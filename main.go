package main

import (
	"encoding/json"
	"flag"
	"time"

	astilectron "github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
	astilog "github.com/asticode/go-astilog"
	"github.com/pkg/errors"
)

var (
	// AppName にはアプリケーション名を格納します.
	AppName string
	// BuiltAt にはアプリケーションが最後に Build された日時を格納します.
	BuiltAt string
	debug   = flag.Bool("d", false, "enables the debug mode")
	w       *astilectron.Window
)

func main() {
	flag.Parse()
	astilog.FlagInit()

	astilog.Debugf("Running app built at %s", BuiltAt)
	if err := bootstrap.Run(bootstrap.Options{
		// Asset コンテンツを取得する
		Asset: Asset,
		// Asset ディレクトリを取得する
		AssetDir: AssetDir,
		// 初期化するときに設定する項目を設定
		// @see https://github.com/asticode/go-astilectron
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDarwinPath:  "resources/kuma.icns",
			AppIconDefaultPath: "resources/kuma.png",
		},
		// Debug モード (flag で付与される, var で default = false としている)
		Debug: *debug,
		// バーティカルメニューの項目を設定
		MenuOptions: []*astilectron.MenuItemOptions{{
			// メニュートップ
			Label: astilectron.PtrStr(AppName),
			// サブメニュー
			SubMenu: []*astilectron.MenuItemOptions{
				{
					Label: astilectron.PtrStr(AppName + "について"),
					OnClick: func(e astilectron.Event) (deleteListener bool) {
						if err := bootstrap.SendMessage(w, "about", htmlAbout(), func(m *bootstrap.MessageIn) {
							// Unmarshal payload
							var s string
							if err := json.Unmarshal(m.Payload, &s); err != nil {
								astilog.Error(errors.Wrap(err, "unmarshaling payload failed"))
								return
							}
							astilog.Infof("About modal has been displayed and payload is %s!", s)
						}); err != nil {
							astilog.Error(errors.Wrap(err, "sending about event failed"))
						}
						return
					},
				},
				// Windowを最小化する
				{Role: astilectron.MenuItemRoleMinimize},
				// Windowを閉じる
				{Role: astilectron.MenuItemRoleClose},
			},
		}},
		// Wait時の動作を設定
		OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			w = ws[0]
			go func() {
				time.Sleep(5 * time.Second)
				if err := bootstrap.SendMessage(w, "check.out.menu", "Don't forget to check out the menu!"); err != nil {
					astilog.Error(errors.Wrap(err, "sending check.out.menu event failed"))
				}
			}()
			return nil
		},
		// Asset をリストアする
		RestoreAssets: RestoreAssets,
		// 画面 (Window) を設定
		Windows: []*bootstrap.Window{{
			Homepage:       "index.html",
			MessageHandler: handleMessages,
			Options: &astilectron.WindowOptions{
				BackgroundColor: astilectron.PtrStr("#333"),
				Center:          astilectron.PtrBool(true),
				Height:          astilectron.PtrInt(700),
				Width:           astilectron.PtrInt(700),
			},
		}},
	}); err != nil {
		astilog.Fatal(errors.Wrap(err, "running bootstrap failed"))
	}
}

func htmlAbout() string {
	return "<b>Kuma</b>と遊ぼう<br>This is a go-astilectron sandbox."
}

// Kuma is the structure of the bear.
type Kuma struct {
	Name string `json:"name"`
}

func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "kuma": // javascript から call される
		// Unmarshal payload
		var name string
		if len(m.Payload) > 0 {
			// Unmarshal payload
			if err = json.Unmarshal(m.Payload, &name); err != nil {
				payload = err.Error()
				return
			}
		}
		if payload, err = kuma(name); err != nil {
			payload = err.Error()
			return
		}
	}
	return
}

func kuma(name string) (k Kuma, err error) {
	k = Kuma{
		Name: name,
	}
	return
}
