```mermaid
graph TD;
    commands;
	config.yaml;
    commands --> |$commandlijn validate| config.yaml;
    commands --> |$commandlijn timetable <halte-id>| Haltes;
    config.yaml --> custom_commands;
	config.yaml --> Haltes;
	config.yaml --> API-keys;
    custom_commands --> Haltes;
