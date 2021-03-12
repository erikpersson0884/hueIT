import React, {useEffect, useState} from "react";
import {postGammaAuth} from "../api/post.GammaAuth";
import {Redirect, Route, useParams} from "react-router";
import {Link} from "react-router-dom";

export const AuthHandler = () => {
    const [redirect, setRedirect] = useState(false)
    const [error, setError] = useState(false)

    useEffect(() => {
        const params = new URLSearchParams(window.location.search)
        postGammaAuth(params.get("code"))
        .then(_ => {
            setRedirect(true)
        })
        .catch(error => {
            console.log("ERR: " + error)
            setError(true)
        })
    })

    return (
        <div>
            { redirect ? (
                <Route>
                    <Redirect to="/"/>
                </Route>
            ) : (
            <div>
                {
                    error ? (
                        <div className="CenterOnPage">
                            <p className="ErrorText">Failed to authenticate with gamma, please contact a system administrator</p>
                            <Link to={"/"}>
                                <button className="BackButton">
                                    Back to main page
                                </button>
                            </Link>
                        </div>
                    ) : (
                        <h1>Authenticating...</h1>
                    )
                }
            </div>
            )
            }
        </div>
    )
}