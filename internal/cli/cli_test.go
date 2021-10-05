package cli_test

import (
	"os"
	"testing"

	cli "github.com/antonio-alexander/go-hello-world/internal/cli"
)

func TestCli(t *testing.T) {
	//range over cases and execute the main with the given
	// inputs, if there's an error output the error, if there's
	// no error, ensure that an error wasn't expected
	chOsSignal := make(chan os.Signal, 1)
	cases := map[string]struct {
		iPwd        string            //present working directroy
		iArgs       []string          //arguments
		iEnv        map[string]string //environmental variables
		iPanic      bool              //panic
		iSignalSend bool              //whether or not to send signal
		iSignal     os.Signal         //os signal
		oErr        string            //error
	}{
		"default": {
			iPwd: "/tmp",
			iEnv: map[string]string{
				"Key": "Value",
			},
			iPanic: false,
		},
		"panic": {
			iPanic: true,
		},
		// "cpu": {
		// 	iArgs: []string{
		// 		"cpu",
		// 	},
		// 	iSignalSend: true,
		// 	iSignal:     os.Interrupt,
		// },
		"memory": {
			iArgs: []string{
				"memory",
			},
			iPanic:      false,
			iSignalSend: true,
			iSignal:     os.Interrupt,
		},
		// "race": {
		// 	iArgs: []string{
		// 		"race",
		// 	},
		// 	iPanic: true,
		// },
		"signal": {
			iArgs: []string{
				"signal",
			},
			iPanic:      false,
			iSignalSend: true,
			iSignal:     os.Interrupt,
		},
		"error": {
			iArgs: []string{
				"error",
			},
			oErr: "Error on purpose",
		},
	}
	for cDesc, c := range cases {
		//if panic is set, check with recovery function,
		// if signal send is set, send the signal and then
		// attempt to run the main, check for errors (expected
		// and unexpected)
		defer func(iPanic bool) {
			if iPanic {
				if r := recover(); r == nil {
					t.Fatalf("Case %s failed, expected panic", cDesc)
				}
			}
		}(c.iPanic)
		if c.iSignalSend {
			chOsSignal <- c.iSignal
		}
		if err := cli.Main(c.iPwd, c.iArgs, c.iEnv, chOsSignal); err != nil {
			if err.Error() != c.oErr {
				t.Fatalf("Case %s failed: %s", cDesc, err)
			}
		} else {
			if c.iPanic {
				t.Fatalf("Case %s failed: panic expected", cDesc)
			}
			if c.oErr != "" {
				t.Fatalf("Case %s failed: error expected", cDesc)
			}
		}
	}
	close(chOsSignal)
}
