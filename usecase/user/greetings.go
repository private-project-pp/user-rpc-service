package user

import context "context"

func (s userService) GreetUser(context.Context, *GreetingRequest) (out *GreetingResponse, err error) {
	return out, nil
}
