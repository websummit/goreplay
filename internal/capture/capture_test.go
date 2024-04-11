package capture

import (
	"testing"
)

func TestSetInterfaces(t *testing.T) {
	listener := &Listener{
		loopIndex: 99999,
	}
	listener.setInterfaces()

	for _, nic := range listener.Interfaces {
		if (len(nic.Addresses)) == 0 {
			t.Errorf("nic %s was captured with 0 addresses", nic.Name)
		}
	}

	if listener.loopIndex == 99999 {
		t.Errorf("loopback nic index was not found")
	}
}

func Test_parseK8sSelector(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name              string
		args              args
		wantNamespace     string
		wantLabelSelector string
		wantFieldSelector string
	}{
		{
			name: "pod",
			args: args{
				"default/pod/test",
			},
			wantNamespace:     "default",
			wantLabelSelector: "",
			wantFieldSelector: "metadata.name=test",
		},
		{
			name: "labelSelectorWithSlash",
			args: args{
				"avenger/labelSelector/app.kubernetes.io/name=avenger:3000",
			},
			wantNamespace:     "avenger",
			wantLabelSelector: "app.kubernetes.io/name=avenger:3000",
			wantFieldSelector: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := parseK8sSelector(tt.args.addr)
			if got != tt.wantNamespace {
				t.Errorf("parseK8sSelector() got = %v, want %v", got, tt.wantNamespace)
			}
			if got1 != tt.wantLabelSelector {
				t.Errorf("parseK8sSelector() got1 = %v, want %v", got1, tt.wantLabelSelector)
			}
			if got2 != tt.wantFieldSelector {
				t.Errorf("parseK8sSelector() got2 = %v, want %v", got2, tt.wantFieldSelector)
			}
		})
	}
}
