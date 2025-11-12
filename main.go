// Copyright 2020 The Nakama Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

var (
	errInternalError  = runtime.NewError("internal server error", 13) // INTERNAL
	errMarshal        = runtime.NewError("cannot marshal type", 13)   // INTERNAL
	errNoInputAllowed = runtime.NewError("no input allowed", 3)       // INVALID_ARGUMENT
	errNoUserIdFound  = runtime.NewError("no user ID in context", 3)  // INVALID_ARGUMENT
	errUnmarshal      = runtime.NewError("cannot unmarshal type", 13) // INTERNAL
)

const (
	rpcIdRewards   = "rewards"
	rpcIdFindMatch = "find_match"
)

// noinspection GoUnusedExportedFunction
func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	//initStart := time.Now()
	//
	//marshaler := &protojson.MarshalOptions{
	//	UseEnumNumbers: true,
	//}
	//unmarshaler := &protojson.UnmarshalOptions{
	//	DiscardUnknown: false,
	//}

	//for adding user
	username := "amir_123"
	password := "12345678"
	email := "test@gmail.com"

	userID, _, _, err := nk.AuthenticateEmail(ctx, email, password, username, true)
	if err != nil {
		logger.WithField("error", err.Error()).Error("❌ Failed to create user")
	} else {
		logger.WithField("user_id", userID).Info("✅ User1 created successfully!")
	}

	username_1 := "amir_124"
	password_1 := "12345687"
	email_1 := "test@gmail.com"

	userID, _, _, err = nk.AuthenticateEmail(ctx, email_1, password_1, username_1, true)
	if err != nil {
		logger.WithField("error", err.Error()).Error("❌ Failed to create user")
	} else {
		logger.WithField("user_id", userID).Info("✅ User2 created successfully!")
	}

	//if err := initializer.RegisterRpc(rpcIdRewards, rpcRewards); err != nil {
	//	return err
	//}
	//
	//if err := initializer.RegisterRpc(rpcIdFindMatch, rpcFindMatch(marshaler, unmarshaler)); err != nil {
	//	return err
	//}
	//
	//if err := initializer.RegisterMatch(moduleName, func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (runtime.Match, error) {
	//	return &MatchHandler{
	//		marshaler:        marshaler,
	//		unmarshaler:      unmarshaler,
	//		tfServingAddress: "http://tf:8501/v1/models/ttt:predict",
	//	}, nil
	//}); err != nil {
	//	return err
	//}
	//
	//if err := registerSessionEvents(db, nk, initializer); err != nil {
	//	return err
	//}

	//logger.Info("Plugin loaded in '%d' msec.", time.Now().Sub(initStart).Milliseconds())
	return nil
}
