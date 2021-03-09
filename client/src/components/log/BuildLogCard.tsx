import React from 'react';
import { IBuildLog } from '../../interfaces/log/IBuildLogInfo';
import { Card } from 'react-bootstrap';

interface IBuildLogCard {
    buildLog: IBuildLog
}

const BuildLogCard: React.FunctionComponent<IBuildLogCard> = ({buildLog}) => {
    return (
        <Card className={'mt-2'}>
            <Card.Body>
                <span className={'font-weight-bold'}>{buildLog.starter}</span> started a
                build
                on <span className={'font-weight-bold'}>{buildLog.project}</span> project
                (<span className={'font-weight-bold'}>{buildLog.module}</span>) for <span
                className={'font-weight-bold'}>{buildLog.target}</span> workgroup,
                <span
                    className={'font-weight-bold'}>{new Date(buildLog.time * 1000).toUTCString()}</span>
            </Card.Body>
        </Card>
    );
};

export default BuildLogCard;
