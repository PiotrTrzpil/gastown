package web

import (
	"testing"

	"github.com/steveyegge/gastown/internal/activity"
)

func TestCalculateWorkStatus(t *testing.T) {
	tests := []struct {
		name          string
		completed     int
		total         int
		activityColor string
		want          string
	}{
		{
			name:          "complete when all done",
			completed:     5,
			total:         5,
			activityColor: activity.ColorGreen,
			want:          "complete",
		},
		{
			name:          "complete overrides activity color",
			completed:     3,
			total:         3,
			activityColor: activity.ColorRed,
			want:          "complete",
		},
		{
			name:          "active when green",
			completed:     2,
			total:         5,
			activityColor: activity.ColorGreen,
			want:          "active",
		},
		{
			name:          "stale when yellow",
			completed:     2,
			total:         5,
			activityColor: activity.ColorYellow,
			want:          "stale",
		},
		{
			name:          "stuck when red",
			completed:     2,
			total:         5,
			activityColor: activity.ColorRed,
			want:          "stuck",
		},
		{
			name:          "waiting when unknown color",
			completed:     2,
			total:         5,
			activityColor: activity.ColorUnknown,
			want:          "waiting",
		},
		{
			name:          "waiting when empty color",
			completed:     0,
			total:         5,
			activityColor: "",
			want:          "waiting",
		},
		{
			name:          "waiting when no work yet",
			completed:     0,
			total:         0,
			activityColor: activity.ColorUnknown,
			want:          "waiting",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateWorkStatus(tt.completed, tt.total, tt.activityColor)
			if got != tt.want {
				t.Errorf("calculateWorkStatus(%d, %d, %q) = %q, want %q",
					tt.completed, tt.total, tt.activityColor, got, tt.want)
			}
		})
	}
}

func TestGetRefineryStatusHint(t *testing.T) {
	// Create a minimal fetcher for testing
	f := &LiveConvoyFetcher{}

	tests := []struct {
		name    string
		mrCount int
		want    string
	}{
		{"idle when no MRs", 0, "Idle - Waiting for MRs"},
		{"singular MR", 1, "Processing 1 MR"},
		{"multiple MRs", 2, "Processing 2 MRs"},
		{"many MRs", 10, "Processing 10 MRs"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := f.getRefineryStatusHint(tt.mrCount)
			if got != tt.want {
				t.Errorf("getRefineryStatusHint(%d) = %q, want %q",
					tt.mrCount, got, tt.want)
			}
		})
	}
}

