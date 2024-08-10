package user

import context "context"

func (s userService) GreetUserLain(context.Context, *GreetingRequest) (out *GreetingResponse, err error) {
	return out, nil
}
