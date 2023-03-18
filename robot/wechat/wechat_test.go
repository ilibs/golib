package wechat

import (
	"os"
	"testing"
)

func TestWechatRobot_Message(t *testing.T) {
	if testing.Short() {
		t.Skip("skip")
	}
	robot := NewRobot()
	robot.SetToken(os.Getenv("WECHAT_ROBOT_TOKEN"))
	err := robot.Message("test")
	if err != nil {
		t.Error(err)
	}
}

func TestWechatRobot_MarkdownMessage(t *testing.T) {
	if testing.Short() {
		t.Skip("skip")
	}
	robot := NewRobot()
	robot.SetToken(os.Getenv("WECHAT_ROBOT_TOKEN"))
	err := robot.MarkdownMessage("## 呵呵\n\n > Hello \n\n")
	if err != nil {
		t.Error(err)
	}
}

func TestWechatRobot_CardMessage(t *testing.T) {
	if testing.Short() {
		t.Skip("skip")
	}
	robot := NewRobot()
	robot.SetToken("16755e18-448c-4ee7-9fbb-e36bdf30f3c3")

	err := robot.CardMessage("要事提醒", "该起床了该起床了该起床了该起床了该起床了该起床了该起床了该起床了该起床了该起床了\n", []map[string]string{
		{
			"title":     "收到提醒",
			"actionURL": "https://www.fifsky.com/",
		},
		{
			"title":     "小睡一下",
			"actionURL": "https://www.fifsky.com/",
		},
	})

	if err != nil {
		t.Error(err)
	}
}
