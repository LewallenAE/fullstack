// Type alias
type HealthResponse = {
    service: string;
    status: "ok" | "unavailable";
};

function buildHealthResponse(serviceName: string, isOperational: boolean): HealthResponse {
    
}