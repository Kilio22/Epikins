import React from "react";

const Excuses: string[] = [
    "rand mon kwad.",
    "attends je mange des coquilettes là.",
    "t'as pas le droit wesh.",
    "t'as surement fait un truc qui fallait pas.",
    "je suis surpris que ca fonctionne.",
    "éssaie de redémarrer ta box pour voir ?",
    "ça fait pas ça chez moi.",
    "dsl pause gouter tmtc.",
    "le franssai ces inportan quan meme.",
    "ptn g perdu ma carte GOLD, PAPAAAAAAAAAAA"
]

const Unauthorized = () => {
    return (
        <div className={"d-flex flex-grow-1 justify-content-center align-items-center text-center"}>
            <h1>{Excuses[Math.floor(Math.random() * Excuses.length)]}</h1>
        </div>
    );
}

export default Unauthorized;