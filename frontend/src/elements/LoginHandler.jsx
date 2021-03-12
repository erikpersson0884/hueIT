import React from "react"

export const LoginHandler = ({loginUrl}) => {
    console.log("LoginUrl: ", loginUrl)
    return (
    <div className="CenterOnPage">
        {
            loginUrl ? (
                <>
                    <p>
                        You need to be logged in to access this page.
                    </p>
                    <button className="LoginButton" onClick={() => {
                        window.location.href = loginUrl;
                    }}>
                        Login
                    </button>
                </>
            ) : (
                <p className="ErrorText">
                    Failed to retrieve login url, please contact a system administrator
                </p>
            )
        }
    </div>
    )
}