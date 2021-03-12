import './App.css';
import React, {useEffect, useState} from "react";
import {AuthHandler} from "./elements/AuthHandler";
import {Main} from "./elements/Main";
import {BrowserRouter} from "react-router-dom";
import {Redirect, Route, Switch} from "react-router";
import {LoginHandler} from "./elements/LoginHandler";
import {initializeSetLoginUrl} from "./utility";

function App() {
    const [loginUrl, setLoginUrl] = useState("")

    useEffect(() => {
        initializeSetLoginUrl(val => {
            setLoginUrl(val)
        })
    }, [])

    return (
    <BrowserRouter>
        {
            loginUrl !== "" && (
                <Route>
                    <Redirect to="/login" />
                </Route>
            )
        }
        <Switch>
            <Route path="/auth/account/callback" component={AuthHandler}/>

            <Route path="/login" component={() => {
                return <LoginHandler loginUrl={loginUrl}/>
            }}/>

            <Route path="/" component={Main}/>
        </Switch>
    </BrowserRouter>
    )
}

export default App;
