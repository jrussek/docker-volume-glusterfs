package main

import (
	"reflect"
	"sync"
	"testing"
)

func Test_newGlusterfsDriver(t *testing.T) {
	type args struct {
		root           string
		defaultServers string
		defaultVolname string
	}
	tests := []struct {
		name    string
		args    args
		want    *glusterfsDriver
		wantErr bool
	}{
		{
			name: "Test newGlusterfsDriver",
			args: args{
				root: "/basedir",
				defaultServers: "localhost",
				defaultVolname: "gluster_volume",
			},
			want: &glusterfsDriver{
				root: "/basedir/volumes",
				statePath: "/basedir/state/gfs-state.json",
				volumes: map[string]*glusterfsVolume{},
				defaultVolname: "gluster_volume",
				defaultServers: "localhost",
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newGlusterfsDriver(tt.args.root, tt.args.defaultServers, tt.args.defaultVolname)
			if (err != nil) != tt.wantErr {
				t.Errorf("newGlusterfsDriver() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newGlusterfsDriver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_glusterfsDriver_saveState(t *testing.T) {
	type fields struct {
		RWMutex        sync.RWMutex
		root           string
		statePath      string
		volumes        map[string]*glusterfsVolume
		defaultVolname string
		defaultServers string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &glusterfsDriver{
				RWMutex:        tt.fields.RWMutex,
				root:           tt.fields.root,
				statePath:      tt.fields.statePath,
				volumes:        tt.fields.volumes,
				defaultVolname: tt.fields.defaultVolname,
				defaultServers: tt.fields.defaultServers,
			}
			d.saveState()
		})
	}
}
