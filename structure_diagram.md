```mermaid
graph TD;
    commandlijn;
    config.yaml;
    commandlijn --> |validate| config.yaml;
    commandlijn --> |init| config.yaml;
    commandlijn --> |timetable [id]| haltes;
    config.yaml --> custom_commands;
	config.yaml --> haltes;
	config.yaml --> API-keys;
