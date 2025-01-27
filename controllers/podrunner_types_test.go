package controllers

import (
	"github.com/evryfs/github-actions-runner-operator/api/v1alpha1"
	"github.com/google/go-github/v33/github"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
	"testing"
	"time"
)

var podList = v1.PodList{
	TypeMeta: metav1.TypeMeta{},
	ListMeta: metav1.ListMeta{},
	Items: []v1.Pod{
		{
			TypeMeta: metav1.TypeMeta{},
			ObjectMeta: metav1.ObjectMeta{
				Name:              "name1",
				CreationTimestamp: metav1.NewTime(time.Now()),
			},
			Spec:   v1.PodSpec{},
			Status: v1.PodStatus{},
		},
		{
			TypeMeta: metav1.TypeMeta{},
			ObjectMeta: metav1.ObjectMeta{
				Name:              "name2",
				CreationTimestamp: metav1.NewTime(time.Now().Add(time.Minute)),
			},
			Spec:   v1.PodSpec{},
			Status: v1.PodStatus{},
		},
	},
}

var runners = []*github.Runner{
	{
		ID:     nil,
		Name:   pointer.StringPtr("name1"),
		OS:     nil,
		Status: nil,
		Busy:   nil,
		Labels: nil,
	},
	{
		ID:     nil,
		Name:   pointer.StringPtr("name2"),
		OS:     nil,
		Status: nil,
		Busy:   nil,
		Labels: nil,
	},
}

func TestPodRunnerPairList(t *testing.T) {
	testCases := []struct {
		podList v1.PodList
		runners []*github.Runner
		inSync  bool
	}{
		{podList, runners, true},
		{v1.PodList{Items: podList.Items[:1]}, runners, false},
		{podList, runners[:1], false},
	}

	for _, tc := range testCases {
		podRunnerPairList := from(&tc.podList, tc.runners)
		assert.Equal(t, podRunnerPairList.inSync(), tc.inSync)
	}
}

func TestSort(t *testing.T) {
	testCases := []struct {
		sortOrder         v1alpha1.SortOrder
		podRunnerPairList podRunnerPairList
		podList           []v1.Pod
	}{
		{v1alpha1.LeastRecent, from(&podList, runners), []v1.Pod{podList.Items[0], podList.Items[1]}},
		{v1alpha1.MostRecent, from(&podList, runners), []v1.Pod{podList.Items[1], podList.Items[0]}},
	}

	for _, tc := range testCases {
		podList := tc.podRunnerPairList.getIdlePods(tc.sortOrder)
		assert.Equal(t, podList, tc.podList)
	}
}
