import {Redirect} from "react-router-dom";
import React from "react";

const Home = () => {
    return (
        <Redirect to={"/projects"}/>
    );
}

export default Home;