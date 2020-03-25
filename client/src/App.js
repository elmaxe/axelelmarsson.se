import React from 'react'
import './App.css'
import Island2019 from './images/island2019.JPG';
import MailIcon from './images/envelope-square-solid.svg';
import GithubIcon from './images/github-square-brands.svg';
import LinkedInIcon from './images/linkedin-brands.svg';
import CVIcon from './images/file-alt-regular.svg';
import PhoneIcon from './images/phone-square-alt-solid.svg';

const App = () => {
    return (
        <div className="Page">
            <div className="AXELRUTA">
                <img src={Island2019} alt="Axel Elmarsson" className="AXEL" title="Nämen hallå! Denna bild är tagen på Island sommaren 2019. Foto: Jesper Sjöholm"/>
            </div>
            <div>
                <h1>Axel Elmarsson</h1>
            </div>
            <div className="Metatext">
                Civilingenjörsstudent i datateknik på KTH.
            </div>
            <div className="Bar">
                    <p>
                        <a
                            href="mailto:axel.elmarsson@gmail.com"
                            title="Mejla mig om du vill, men föredrar brev! /s"
                            >
                            <img src={MailIcon} alt="E-mail" className="Images"/>
                        </a>
                        <a
                            href="https://github.com/elmaxe"
                            title="Min Github :O"
                            rel="noopener noreferrer" target="_blank"
                            >
                            <img src={GithubIcon} alt="Github logo" className="Images"/>
                        </a>
                        <a
                            href="/public/CV_AXEL_ELMARSSON.pdf"
                            title="Mitt CV..."
                            rel="noopener noreferrer" target="_blank"
                            >
                            <img src={CVIcon} alt="CV" className="Images CV"/>
                        </a>
                        <a
                            href="https://www.linkedin.com/in/axel-elmarsson-453a93159/"
                            title="Lägg till mig på LinkedIn!"
                            rel="noopener noreferrer" target="_blank"
                            >
                            <img src={LinkedInIcon} alt="LinkedIn logo" className="Images"/>
                        </a>
                        <a 
                            href="tel:+46767901038"
                            title="Ring mig inte pls!"
                            >
                            <img src={PhoneIcon} alt="Phone" className="Images"/>
                        </a>

                    </p>
            </div>
            <div className="Fötter">
                Skrivet i React.js av Axel Elmarsson 2020
            </div>
        </div>
    )
}

export default App
