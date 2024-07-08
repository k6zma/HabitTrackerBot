# Habit Tracking Telegram Bot in Go

Welcome! This project is a simple habit-tracking Telegram bot written in Go, designed to help users manage and monitor their habits without using a database.

## Overview

The bot allows users to:
- Add new habits
- Delete existing habits
- Mark habits as completed
- Unmark completed habits
- View the list of all habits
- See the completion statistics of their habits

## Requirements

1. The bot must be written in Go.
2. Data should be stored in memory using structures like maps and slices.
3. The bot should support the following commands:

   - `/start` - Introduction and explanation of the bot's functionality.
   - `/help` - List of all available commands and their descriptions.
   - `/add_habit <name>` - Add a new habit.
   - `/delete_habit <name>` - Delete a habit.
   - `/list_habits` - Show the list of all habits.
   - `/mark_habit <name>` - Mark a habit as completed.
   - `/unmark_habit <name>` - Unmark a completed habit.
   - `/stats` - Show the completion statistics of all habits.