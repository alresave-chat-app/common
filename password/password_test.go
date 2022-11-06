package password

import "testing"

func TestValidatePassword(t *testing.T) {
	hashed, _ := HashPassword("password")
	type args struct {
		hashed   string
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ValidatePassword with right password should validate",
			args: args{
				hashed:   hashed,
				password: "password",
			},
			want: true,
		},
		{
			name: "ValidatePassword with wrong password should not validate",
			args: args{
				hashed:   hashed,
				password: "nopassword",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidatePassword(tt.args.hashed, tt.args.password); got != tt.want {
				t.Errorf("ValidatePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
