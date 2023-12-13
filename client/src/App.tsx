import React from 'react';
import logo from './logo.svg';
import './App.css';
import { AuditLogViewer, Audit } from '@pangeacyber/react-mui-audit-log-viewer';
import { Box } from '@mui/material';

const handleSearch = (request: any): Promise<Audit.SearchResponse> => {
  return fetch("http://localhost:4000/audit/search", {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify(request),
  }).then(async (res) => {
    const data = await res.json();
    console.log(data);
    return data?.result;
  });
};

const handlePageChange = (request: any): Promise<Audit.ResultResponse> => {
  return fetch("http://localhost:4000/audit/results", {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify(request),
  }).then(async (res) => {
    const data = await res.json();
    return data?.result;
  });
};

const handleFetchRoot = (request: any): Promise<Audit.RootResponse> => {
  return fetch("http://localhost:4000/audit/root", {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify(request),
  }).then(async (res) => {
    const data = await res.json();
    return data?.result;
  });
};

function App() {
  return (
    <>
    <div className="App">
      <header className="App-header">
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        </header>
    </div>
        <Box sx={{ margin: "8px", width: "calc(100% - 64px)", padding: "32px" }}>
        <AuditLogViewer
          onSearch={handleSearch}
          onPageChange={handlePageChange}
          verificationOptions={{
            onFetchRoot: handleFetchRoot
          }} 
        />
        </Box>
    </>
  );
}

export default App;
