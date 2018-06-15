// Automatically generated by pegomock. DO NOT EDIT!
// Source: github.com/runatlantis/atlantis/server/events (interfaces: PullRequestOperator)

package mocks

import (
	"reflect"

	pegomock "github.com/petergtz/pegomock"
	events "github.com/runatlantis/atlantis/server/events"
)

type MockPullRequestOperator struct {
	fail func(message string, callerSkip ...int)
}

func NewMockPullRequestOperator() *MockPullRequestOperator {
	return &MockPullRequestOperator{fail: pegomock.GlobalFailHandler}
}

func (mock *MockPullRequestOperator) Autoplan(ctx *events.CommandContext) events.CommandResponse {
	params := []pegomock.Param{ctx}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Autoplan", params, []reflect.Type{reflect.TypeOf((*events.CommandResponse)(nil)).Elem()})
	var ret0 events.CommandResponse
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(events.CommandResponse)
		}
	}
	return ret0
}

func (mock *MockPullRequestOperator) PlanViaComment(ctx *events.CommandContext) events.CommandResponse {
	params := []pegomock.Param{ctx}
	result := pegomock.GetGenericMockFrom(mock).Invoke("PlanViaComment", params, []reflect.Type{reflect.TypeOf((*events.CommandResponse)(nil)).Elem()})
	var ret0 events.CommandResponse
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(events.CommandResponse)
		}
	}
	return ret0
}

func (mock *MockPullRequestOperator) ApplyViaComment(ctx *events.CommandContext) events.CommandResponse {
	params := []pegomock.Param{ctx}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ApplyViaComment", params, []reflect.Type{reflect.TypeOf((*events.CommandResponse)(nil)).Elem()})
	var ret0 events.CommandResponse
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(events.CommandResponse)
		}
	}
	return ret0
}

func (mock *MockPullRequestOperator) VerifyWasCalledOnce() *VerifierPullRequestOperator {
	return &VerifierPullRequestOperator{mock, pegomock.Times(1), nil}
}

func (mock *MockPullRequestOperator) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierPullRequestOperator {
	return &VerifierPullRequestOperator{mock, invocationCountMatcher, nil}
}

func (mock *MockPullRequestOperator) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierPullRequestOperator {
	return &VerifierPullRequestOperator{mock, invocationCountMatcher, inOrderContext}
}

type VerifierPullRequestOperator struct {
	mock                   *MockPullRequestOperator
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierPullRequestOperator) Autoplan(ctx *events.CommandContext) *PullRequestOperator_Autoplan_OngoingVerification {
	params := []pegomock.Param{ctx}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Autoplan", params)
	return &PullRequestOperator_Autoplan_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type PullRequestOperator_Autoplan_OngoingVerification struct {
	mock              *MockPullRequestOperator
	methodInvocations []pegomock.MethodInvocation
}

func (c *PullRequestOperator_Autoplan_OngoingVerification) GetCapturedArguments() *events.CommandContext {
	ctx := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1]
}

func (c *PullRequestOperator_Autoplan_OngoingVerification) GetAllCapturedArguments() (_param0 []*events.CommandContext) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*events.CommandContext, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(*events.CommandContext)
		}
	}
	return
}

func (verifier *VerifierPullRequestOperator) PlanViaComment(ctx *events.CommandContext) *PullRequestOperator_PlanViaComment_OngoingVerification {
	params := []pegomock.Param{ctx}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "PlanViaComment", params)
	return &PullRequestOperator_PlanViaComment_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type PullRequestOperator_PlanViaComment_OngoingVerification struct {
	mock              *MockPullRequestOperator
	methodInvocations []pegomock.MethodInvocation
}

func (c *PullRequestOperator_PlanViaComment_OngoingVerification) GetCapturedArguments() *events.CommandContext {
	ctx := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1]
}

func (c *PullRequestOperator_PlanViaComment_OngoingVerification) GetAllCapturedArguments() (_param0 []*events.CommandContext) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*events.CommandContext, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(*events.CommandContext)
		}
	}
	return
}

func (verifier *VerifierPullRequestOperator) ApplyViaComment(ctx *events.CommandContext) *PullRequestOperator_ApplyViaComment_OngoingVerification {
	params := []pegomock.Param{ctx}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ApplyViaComment", params)
	return &PullRequestOperator_ApplyViaComment_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type PullRequestOperator_ApplyViaComment_OngoingVerification struct {
	mock              *MockPullRequestOperator
	methodInvocations []pegomock.MethodInvocation
}

func (c *PullRequestOperator_ApplyViaComment_OngoingVerification) GetCapturedArguments() *events.CommandContext {
	ctx := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1]
}

func (c *PullRequestOperator_ApplyViaComment_OngoingVerification) GetAllCapturedArguments() (_param0 []*events.CommandContext) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*events.CommandContext, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(*events.CommandContext)
		}
	}
	return
}
