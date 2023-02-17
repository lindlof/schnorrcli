package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	btcec "github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "pubkey",
				Usage: "generate schnorr pubkey from private key",
				Action: func(cCtx *cli.Context) error {
					b, err := base64.StdEncoding.DecodeString(cCtx.Args().First())
					if err != nil {
						panic(err)
					}
					priv, _ := btcec.PrivKeyFromBytes(b)

					pub := schnorr.SerializePubKey(priv.PubKey())
					fmt.Printf("%s\n", base64.StdEncoding.EncodeToString(pub))
					return nil
				},
			},
			{
				Name:  "signb64",
				Usage: "sign a base64 document",
				Action: func(cCtx *cli.Context) error {
					b, err := base64.StdEncoding.DecodeString(cCtx.Args().First())
					if err != nil {
						panic(err)
					}
					priv, _ := btcec.PrivKeyFromBytes(b)

					doc, err := base64.StdEncoding.DecodeString(cCtx.Args().Get(1))
					if err != nil {
						panic(err)
					}
					h := sha256.New()
					h.Write(doc)

					signature, _ := schnorr.Sign(priv, h.Sum(nil))
					fmt.Printf("%s\n", base64.StdEncoding.EncodeToString(signature.Serialize()))
					return nil
				},
			},
			{
				Name:  "sign",
				Usage: "sign a string document",
				Action: func(cCtx *cli.Context) error {
					b, err := base64.StdEncoding.DecodeString(cCtx.Args().First())
					if err != nil {
						panic(err)
					}
					priv, _ := btcec.PrivKeyFromBytes(b)

					h := sha256.New()
					h.Write([]byte(cCtx.Args().Get(1)))

					signature, _ := schnorr.Sign(priv, h.Sum(nil))
					fmt.Printf("%s\n", base64.StdEncoding.EncodeToString(signature.Serialize()))
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
