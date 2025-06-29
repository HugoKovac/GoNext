import { useEffect, useState } from "react";

function Profile() {
    const [error, setError] = useState("");
    const [email, setEmail] = useState("");
    const [oldPassword, setOldPassword] = useState("");
    const [newPassword, setNewPassword] = useState("");

    useEffect(() => {
        fetch(`${import.meta.env.VITE_API_URL}/api/users/me`, {
            credentials: "include",
        })
            .then((res) => {
                if (res.ok) {
                    return res.json();
                }
                throw new Error("Failed to fetch user");
            })
            .then((data) => {
                setEmail(data.email);
            })
            .catch((err) => {
                console.error(err);
            });
    }, []);

    const handleSubmit = async (e) => {
        e.preventDefault();
        fetch(`${import.meta.env.VITE_API_URL}/api/users/me`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ email, oldPassword, newPassword }),
        }).then((res) => {
            console.log(res);
        });
    };
    return (
        <div className="flex flex-col h-screen w-screen justify-center">
            <div className="card w-96 bg-base-100 shadow-sm mx-auto border-2 max-w-6/7 py-6">
                <form className="card-body flex flex-col" onSubmit={handleSubmit}>
                    <div className="flex justify-around">
                        <h2 className="text-3xl font-bold">Profile</h2>
                    </div>
                    <div className="my-6 flex flex-col">
                        <label className="input validator my-2">
                            <svg
                                className="h-[1em] opacity-50"
                                xmlns="http://www.w3.org/2000/svg"
                                viewBox="0 0 24 24"
                            >
                                <g
                                    strokeLinejoin="round"
                                    strokeLinecap="round"
                                    strokeWidth="2.5"
                                    fill="none"
                                    stroke="currentColor"
                                >
                                    <rect width="20" height="16" x="2" y="4" rx="2"></rect>
                                    <path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"></path>
                                </g>
                            </svg>
                            <input
                                type="email"
                                placeholder="mail@site.com"
                                required
                                name="email"
                                value={email}
                                onChange={(e) => setEmail(e.target.value)}
                            />
                        </label>
                        <div className="validator-hint hidden">
                            Enter valid email address
                        </div>

                        <label className="input validator my-2">
                            <svg
                                className="h-[1em] opacity-50"
                                xmlns="http://www.w3.org/2000/svg"
                                viewBox="0 0 24 24"
                            >
                                <g
                                    strokeLinejoin="round"
                                    strokeLinecap="round"
                                    strokeWidth="2.5"
                                    fill="none"
                                    stroke="currentColor"
                                >
                                    <path d="M2.586 17.414A2 2 0 0 0 2 18.828V21a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1h1a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1h.172a2 2 0 0 0 1.414-.586l.814-.814a6.5 6.5 0 1 0-4-4z"></path>
                                    <circle
                                        cx="16.5"
                                        cy="7.5"
                                        r=".5"
                                        fill="currentColor"
                                    ></circle>
                                </g>
                            </svg>
                            <input
                                type="password"
                                name="oldPassword"
                                placeholder="Old Password"
                                minLength="12"
                                pattern="^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*]).{12,}$"
                                title="Password must be at least 12 characters long and include at least one uppercase letter, one lowercase letter, one number, and one special character"
                                value={oldPassword}
                                onChange={(e) => setOldPassword(e.target.value)}
                            />
                        </label>
                        <p className="validator-hint hidden">
                            Must be more than 8 characters, including
                            <br />
                            At least one number <br />
                            At least one lowercase letter <br />
                            At least one uppercase letter
                        </p>
                        <label className="input validator my-2">
                            <svg
                                className="h-[1em] opacity-50"
                                xmlns="http://www.w3.org/2000/svg"
                                viewBox="0 0 24 24"
                            >
                                <g
                                    strokeLinejoin="round"
                                    strokeLinecap="round"
                                    strokeWidth="2.5"
                                    fill="none"
                                    stroke="currentColor"
                                >
                                    <path d="M2.586 17.414A2 2 0 0 0 2 18.828V21a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1h1a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1h.172a2 2 0 0 0 1.414-.586l.814-.814a6.5 6.5 0 1 0-4-4z"></path>
                                    <circle
                                        cx="16.5"
                                        cy="7.5"
                                        r=".5"
                                        fill="currentColor"
                                    ></circle>
                                </g>
                            </svg>
                            <input
                                type="password"
                                name="newPassword"
                                placeholder="New Password"
                                minLength="12"
                                pattern="^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*]).{12,}$"
                                title="Password must be at least 12 characters long and include at least one uppercase letter, one lowercase letter, one number, and one special character"
                                value={newPassword}
                                onChange={(e) => setNewPassword(e.target.value)}
                            />
                        </label>

                        {error && <p className="text-error">{error}</p>}

                    </div>

                    <div className="flex flex-col justify-around">
                        <button className="btn btn-primary my-1">Update Profile</button>
                    </div>
                </form>
            </div>
        </div>
    );
}

export default Profile;