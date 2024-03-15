import {isRouteErrorResponse, json, useRouteError} from "react-router-dom";

export async function loader() {
    console.log("Throwing json error from loader...")
    throw json(
        {
            sorry: "You have issues.",
            hrEmail: "jorge.xyz@outlook.com",
        },
        {status: 401}
    );
}

export function ErrorBoundary() {
    const error = useRouteError();
    console.log(error);

    if (isRouteErrorResponse(error)) {
        // the response json is automatically parsed to
        // `error.data`, you also have access to the status
        return (
            <div>
                <h1>{error.status}</h1>
                <h2>{error.data.sorry}</h2>
                <p>
                    Go ahead and email {error.data.hrEmail} if you
                    feel like this is a mistake.
                </p>
            </div>
        );
    }

    // rethrow to let the parent error boundary handle it
    // when it's not a special case for this route
    return <div>Dang!</div>;
}