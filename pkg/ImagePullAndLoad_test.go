package pkg

import "testing"

func TestImagePullAndTag(t *testing.T) {
	node := "master"
	if err:= ImagePullAndTag(node); err != nil {
		t.Errorf("ImagePullAndTag testing error, err is: %s", err)
	}

	node = "worker"
	if err:= ImagePullAndTag(node); err != nil {
		t.Errorf("ImagePullAndTag testing error, err is: %s", err)
	}
}
