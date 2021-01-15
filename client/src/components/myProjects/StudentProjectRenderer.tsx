import React from 'react';
import { Button } from 'react-bootstrap';
import { OnStudentProjectClick } from '../../interfaces/Functions';
import { IStudentJob } from '../../interfaces/myProjects/IStudentJob';

interface IStudentProjectRendererProps {
    job: IStudentJob,
    onStudentProjectClick: OnStudentProjectClick
}

const cssColorArray: string[] = [
    'red', 'tomato', 'orange', 'green'
];

const StudentProjectRenderer: React.FunctionComponent<IStudentProjectRendererProps> = ({
                                                                                           job,
                                                                                           onStudentProjectClick
                                                                                       }) => {
    return (
        <Button variant={'outline-primary'}
                className={'m-1 text-left d-flex align-items-center'}
                disabled={job.mongoWorkgroupData.remainingBuilds === 0}
                block={true}
                onClick={(() => onStudentProjectClick(job))}>
            <i className={'fas fa-user-friends mr-1'}/> {job.project.name}
            <div className={'ml-auto mr-1'}>
                [
                {
                    job.mongoWorkgroupData.remainingBuilds >= 4 ?
                        <span className={'font-weight-bold remaining-builds-green'}>
                                {job.mongoWorkgroupData.remainingBuilds}
                        </span>
                        :
                        <span className={'font-weight-bold remaining-builds-' +
                        cssColorArray[job.mongoWorkgroupData.remainingBuilds]}>
                                {job.mongoWorkgroupData.remainingBuilds}
                        </span>
                }
                ]
            </div>
        </Button>
    );
};

export default StudentProjectRenderer;
