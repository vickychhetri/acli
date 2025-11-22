acli is a lightweight CLI tool for recording your daily work logs.
			Each task you add is automatically saved into a date-wise CSV file, making it
			easy to track what you worked on throughout the day.

			Use acli to add tasks, list tasks for a specific date, and maintain a clean
			timeline of your work activities. Ideal for developers, professionals, or anyone
			who wants a fast and minimal personal logging system.

Location: $HOME/.local/share/acli/logs/

Install: sudo dpkg -i acli_1.0.0_amd64.deb


Use

**Add**
acli add "Fixed login bug" --priority high --category backend

--priority low|medium|high
--category string
--date YYYY-MM-DD

acli add "Client meeting" --priority medium --category meeting --date 2025-01-10

**LIST**
acli list

List logs for a specific date:
acli list --date 2025-01-10

**Weekly**
acli weekly
