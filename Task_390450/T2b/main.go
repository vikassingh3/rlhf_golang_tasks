// Import localForage for secure storage
import localForage from "localforage";

// Set the token in localForage
async function setToken(token) {
  try {
    await localForage.setItem("jwtToken", token);
  } catch (error) {
    console.error("Error setting token:", error);
  }
}

// Get the token from localForage
async function getToken() {
  try {
    const token = await localForage.getItem("jwtToken");
    return token;
  } catch (error) {
    console.error("Error getting token:", error);
    return null;
  }
}

// Remove the token from localForage
async function removeToken() {
  try {
    await localForage.removeItem("jwtToken");
  } catch (error) {
    console.error("Error removing token:", error);
  }
}

// Example usage:
async function login() {
  const username = "user";
  const password = "password";
  // Call the Login function from the WebAssembly module
  const loginRequest = exports.protobuf.LoginRequest.create({ username, password });
  const loginResponse = exports.protobuf.LoginResponse.decode(exports.Login(loginRequest));
  const token = loginResponse.token;
  // Store the token securely in localForage
  await setToken(token);
}

async function validateToken() {
  const token = await getToken();
  if (!token) {
    console.log("No token found.");
    return;
  }