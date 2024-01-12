package headless

import (
	"context"
	"github.com/chromedp/chromedp"
	"time"
)

func NewChromedp(headless bool, proxy string) context.Context {
	opts := append(
		// 以默认配置的数组为基础，覆写 headless 参数
		// 当然也可以根据自己的需要进行修改，这个 flag 是浏览器的设置
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoSandbox,
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("headless", headless), // 显示界面
		chromedp.Flag("disable-web-security", true),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.Flag("enable-automation", false),                       // 防止监测 webdriver
		chromedp.Flag("disable-blink-features", "AutomationControlled"), // 禁用 blink 特征，绕过了加速乐检测
	)
	if proxy != "" {
		opts = append(opts, chromedp.ProxyServer(proxy))
	}
	ctx, _ := chromedp.NewExecAllocator(
		context.Background(),
		opts...,
	)
	ctx, _ = chromedp.NewContext(
		ctx,
	)

	return ctx
}

func NewChromedpWithTimeout(headless bool, proxy string, timeout time.Duration) (context.Context, context.CancelFunc) {
	ctx := NewChromedp(headless, proxy)
	ctx, cancel := context.WithTimeout(ctx, timeout)
	return ctx, cancel
}
