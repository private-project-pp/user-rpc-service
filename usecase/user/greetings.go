package user

import context "context"

func (s userService) GreetUser(context.Context, *GreetingRequest) (out *GreetingResponse, err error) {
	out = &GreetingResponse{
		GreetingMessage: "Halo gan",
	}
	return out, nil
}
