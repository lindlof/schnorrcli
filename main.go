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
						return err
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
					doc, err := base64.StdEncoding.DecodeString(cCtx.Args().Get(1))
					if err != nil {
						return err
					}
					return sign(cCtx, doc)
				},
			},
			{
				Name:  "sign",
				Usage: "sign a string document",
				Action: func(cCtx *cli.Context) error {
					return sign(cCtx, []byte(cCtx.Args().Get(1)))
				},
			},
			{
				Name:  "verify",
				Usage: "verify a schnorr signature",
				Action: func(cCtx *cli.Context) error {
					return verify(cCtx, []byte(cCtx.Args().Get(2)))
				},
			},
			{
				Name:  "verifyb64",
				Usage: "verify a schnorr signature",
				Action: func(cCtx *cli.Context) error {
					doc, err := base64.StdEncoding.DecodeString(cCtx.Args().Get(2))
					if err != nil {
						return err
					}
					return verify(cCtx, doc)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func sign(cCtx *cli.Context, doc []byte) error {
	b, err := base64.StdEncoding.DecodeString(cCtx.Args().First())
	if err != nil {
		return err
	}
	priv, _ := btcec.PrivKeyFromBytes(b)

	h := sha256.New()
	h.Write(doc)

	signature, _ := schnorr.Sign(priv, h.Sum(nil))
	fmt.Printf("%s\n", base64.StdEncoding.EncodeToString(signature.Serialize()))
	return nil
}

func verify(cCtx *cli.Context, doc []byte) error {
	b, err := base64.StdEncoding.DecodeString(cCtx.Args().First())
	if err != nil {
		return err
	}

	publicKey, err := schnorr.ParsePubKey(b)
	if err != nil {
		return err
	}

	sigBytes, err := base64.StdEncoding.DecodeString(cCtx.Args().Get(1))
	if err != nil {
		return err
	}

	signature, err := schnorr.ParseSignature(sigBytes)
	if err != nil {
		return err
	}

	h := sha256.New()
	h.Write(doc)

	if !signature.Verify(h.Sum(nil), publicKey) {
		fmt.Printf("signature does not verify\n")
		os.Exit(1)
	}

	fmt.Printf("signature verifies\n")
	return nil
}
