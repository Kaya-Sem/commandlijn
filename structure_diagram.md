```mermaid
graph TD;
    commandlijn;
    config.yaml;
    commandlijn --> |$commandlijn validate| config.yaml;
    commandlijn --> |$commandlijn timetable <id>| haltes;
    config.yaml --> custom_commands;
	config.yaml --> haltes;
	config.yaml --> API-keys;
    custom_commands --> haltes;
