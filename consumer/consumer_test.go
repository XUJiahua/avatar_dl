package consumer

import "testing"

func TestConsumer_Do(t *testing.T) {
	err := New("").Do("http://thirdwx.qlogo.cn/mmopen/a4d6EO3X4ZVDGic8Y3ep5gLdFiclicUAuAcqMUKUGUSlhiarhHibDaf0Soeu6MONRAFj8CgZQicGiaovJUmWHj0Q8o2BedibDaHQ8Pgm/132")
	if err != nil {
		t.Log(err)
	}
}

func Test_genFilename(t *testing.T) {
	t.Log(genFilename("http://thirdwx.qlogo.cn/mmopen/a4d6EO3X4ZVFRwRlWLMDRLU8icwnNIEv7VD5YbhALn5enOjoPxRBeJtXkQ321cSu6JpjamkQkpsTBY3o1QXPmVkoaGFGLucyia/132"))
}
