import React from 'react';
import ReactDOM from 'react-dom';
import './css/index.css';
import './css/jobsRenderer.css';
import './css/footer.css';
import './css/users.css';
import './css/projectsManagement.css';
import './css/myProjects.css';
import './css/buildLog.css';
import * as serviceWorker from './serviceWorker';
import App from './components/App';

ReactDOM.render(
    <React.StrictMode>
        <App/>
    </React.StrictMode>,
    document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
