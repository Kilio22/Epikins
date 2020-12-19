import React from "react";

const Legend = () => {
    return (
        <div className={"ml-1 mt-2"}>
            <p>Legend:</p>
            <ul>
                <li><i className="fas fa-clock"/> in queue / currently building</li>
            </ul>
        </div>
    )
};

export default Legend;