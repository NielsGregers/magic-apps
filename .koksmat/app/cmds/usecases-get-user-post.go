// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: Get  User
---
*/
package cmds

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/magicbutton/magic-apps/execution"
	"github.com/magicbutton/magic-apps/schemas"
	"github.com/magicbutton/magic-apps/utils"
)

func UsecasesGetUserPost(ctx context.Context, args []string) (*schemas.User, error) {

	result, pwsherr := execution.ExecutePowerShell("john", "*", "magic-apps", "05-usecases", "10-get-user.ps1", "", "-userid", args[0])
	if pwsherr != nil {
		return nil, pwsherr
	}

	resultingFile := path.Join(utils.WorkDir("magic-apps"), "user.json")
	data, err := os.ReadFile(resultingFile)
	if err != nil {
		return nil, err
	}
	resultObject := schemas.User{}
	err = json.Unmarshal(data, &resultObject)
	if utils.Output == "json" {
		fmt.Println(string(data))
	}
	utils.PrintSkip2FirstAnd2LastLines(string(result))

	return nil, nil

}