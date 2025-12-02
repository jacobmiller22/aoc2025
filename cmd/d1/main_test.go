package main

import "testing"

func TestRotate(t *testing.T) {

	ringsize := 100
	testCases := []struct {
		name       string
		positiion  int
		direction  Rotation
		distance   int
		wantClicks int
		wantNewPos int
	}{
		{
			name:       "Simple Left",
			positiion:  50,
			direction:  RotationLeft,
			distance:   5,
			wantClicks: 0,
			wantNewPos: 45,
		},
		{
			name:       "Left, to zero",
			positiion:  25,
			direction:  RotationLeft,
			distance:   25,
			wantClicks: 1,
			wantNewPos: 0,
		},
		{
			name:       "Left, to zero from middle",
			positiion:  50,
			direction:  RotationLeft,
			distance:   50,
			wantClicks: 1,
			wantNewPos: 0,
		},
		{
			name:       "Left, to zero from middle, multiple clicks",
			positiion:  50,
			direction:  RotationLeft,
			distance:   1000,
			wantClicks: 10,
			wantNewPos: 50,
		},
		{
			name:       "Simple Right",
			positiion:  50,
			direction:  RotationRight,
			distance:   5,
			wantClicks: 0,
			wantNewPos: 55,
		},
		{
			name:       "Right, to zero",
			positiion:  75,
			direction:  RotationRight,
			distance:   25,
			wantClicks: 1,
			wantNewPos: 0,
		},
		{
			name:       "Right, to zero from middle",
			positiion:  50,
			direction:  RotationRight,
			distance:   50,
			wantClicks: 1,
			wantNewPos: 0,
		},
		{
			name:       "Right, to zero from middle, multiple clicks",
			positiion:  50,
			direction:  RotationRight,
			distance:   1000,
			wantClicks: 10,
			wantNewPos: 50,
		},
		{
			name:       "Right, from zero, full ring",
			positiion:  0,
			direction:  RotationRight,
			distance:   100,
			wantClicks: 1,
			wantNewPos: 0,
		},
		{
			name:       "Right, from zero, partial",
			positiion:  0,
			direction:  RotationRight,
			distance:   5,
			wantClicks: 0,
			wantNewPos: 5,
		},
		{
			name:       "Left 5, from zero to 95",
			positiion:  0,
			direction:  RotationLeft,
			distance:   5,
			wantClicks: 0,
			wantNewPos: 95,
		},
		{
			name:       "Right 100, from zero to 0",
			positiion:  0,
			direction:  RotationLeft,
			distance:   100,
			wantClicks: 1,
			wantNewPos: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			gotNewPos, gotClicks := rotate(ringsize, tc.positiion, tc.direction, tc.distance)

			if gotNewPos != tc.wantNewPos {
				t.Errorf("new position mismatch, (+got, -want), +%d -%d", gotNewPos, tc.wantNewPos)
			}

			if gotClicks != tc.wantClicks {
				t.Errorf("new clicks mismatch, got %d, want %d", gotClicks, tc.wantClicks)
			}
		})
	}
}
