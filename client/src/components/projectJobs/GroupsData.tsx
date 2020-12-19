import React from "react";
import {IGroupData} from "../../interfaces/IGroupData";
import {OnCheckboxChange, OnJobClick} from "../../interfaces/Functions";
import {Button, Col, Row} from "react-bootstrap";
import {appInitialContext} from "../../interfaces/IAppContext";

interface IStudentRemainingBuildsProps {
    groupData: IGroupData
}

interface IGroupDataProps {
    groupData: IGroupData,
    selectedJobs: string[],
    onCheckboxChange: OnCheckboxChange,
    onJobClick: OnJobClick
}

interface IGroupsDataProps {
    groupsData: IGroupData[],
    selectedJobs: string[],
    onCheckboxChange: OnCheckboxChange,
    onJobClick: OnJobClick
}

const cssColorArray: string[] = [
    "red", "tomato", "orange", "green", "green", "green"
]

const GroupMasterRemainingBuilds: React.FunctionComponent<IStudentRemainingBuildsProps> = ({groupData}) => {
    return (
        <appInitialContext.Consumer>
            {context => (
                <div className={"ml-auto mr-1"}>
                    [
                    {
                        context.fuMode ?
                            <span className={"font-weight-bold remaining-builds-green"}>
                                ∞
                            </span>
                            :
                            <span className={"font-weight-bold remaining-builds-" +
                            cssColorArray[groupData.mongoGroupData.remainingBuilds % cssColorArray.length]}>
                                {groupData.mongoGroupData.remainingBuilds}
                            </span>
                    }
                    ]
                </div>
            )}
        </appInitialContext.Consumer>
    );
};

const GroupData: React.FunctionComponent<IGroupDataProps> = ({groupData, selectedJobs, onCheckboxChange, onJobClick}) => {
    return (
        <Button variant={"outline-primary"}
                className={"m-1 text-left d-flex align-items-center"}
                block={true}
                onClick={(event => onJobClick(event, groupData.groupJob.job.url))}>
            <input
                className={"jobs-checkbox mr-2"}
                type={"checkbox"}
                checked={selectedJobs.includes(groupData.groupJob.job.name)}
                onChange={(event => onCheckboxChange(event.target.checked, groupData))}
                disabled={!groupData.mongoGroupData.remainingBuilds}
            />
            {' '}
            <i className={"fas fa-folder mr-1"}/> {groupData.groupJob.job.name}
            <GroupMasterRemainingBuilds groupData={groupData}/>
            {
                (groupData.groupJob.jobInfos.inQueue || groupData.groupJob.jobInfos.lastBuild.buildInfos.building) &&
                <i className="fas fa-clock fa-color-clock jobs-clock-icon"/>
            }
        </Button>
    )
}

const GroupsData: React.FunctionComponent<IGroupsDataProps> = ({groupsData, selectedJobs, onCheckboxChange, onJobClick}) => {
    return (
        <div className={"mt-3"}>
            {
                groupsData.length === 0 ?
                    <h2 className={"text-center"}>No jobs to display</h2>
                    :
                    <Row>
                        {
                            groupsData.map((job, id) => {
                                return (
                                    <Col md={4} key={id}>
                                        <GroupData groupData={job} selectedJobs={selectedJobs}
                                                   onCheckboxChange={onCheckboxChange}
                                                   onJobClick={onJobClick}/>
                                    </Col>
                                );
                            })
                        }
                    </Row>
            }
        </div>
    );
};

export default GroupsData;