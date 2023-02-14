// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package tui

import "github.com/charmbracelet/lipgloss"

var (
	listMainStyle  = lipgloss.NewStyle().Margin(0, 2, 1)
	askMainStyle   = lipgloss.NewStyle().Margin(0, 2, 1, 4) // A wider margin is used to align with lists.
	titleStyle     = lipgloss.NewStyle().Bold(true)
	descStyle      = lipgloss.NewStyle()
	tableMainStyle = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240"))
)
