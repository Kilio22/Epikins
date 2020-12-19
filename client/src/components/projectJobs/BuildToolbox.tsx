import {ButtonGroup, Spinner} from "react-bootstrap";
import BuildDropdown from "./BuildDropdown";
import React from "react";
import {OnBuildClick} from "../../interfaces/Functions";

interface IBuildToolboxProps {
    selectedJobs: string[],
    isBuilding: boolean,
    onBuildClick: OnBuildClick,
    onGlobalBuildClick: OnBuildClick
}

const BuildToolbox: React.FunctionComponent<IBuildToolboxProps> = ({selectedJobs, isBuilding, onBuildClick, onGlobalBuildClick}) => {
    return (
        <div className={"build-toolbox"}>
            <ButtonGroup>
                <BuildDropdown title={"Build"}
                               disabled={selectedJobs.length === 0 || isBuilding}
                               onBuildClick={onBuildClick}/>
                <BuildDropdown title={"Global build"} disabled={isBuilding}
                               onBuildClick={onGlobalBuildClick}/>
            </ButtonGroup>
            {
                isBuilding &&
                <span>
                    <Spinner animation={"border"} variant={"primary"} style={{width: '20px', height: '20px'}}/> starting builds...
                </span>
            }
        </div>
    );
};

export default BuildToolbox;