import qs from 'npm:qs';

const port = 3000;

const handler = async (request: Request): Promise<Response> => {
    let body = await request.text();
    try {
        body = JSON.parse(body);
    } catch {
        /* nothing sauce */
    }

    const response = {
        url: request.url,
        method: request.method,
        query: qs.parse(new URL(request.url).search.substring(1)),
        headers: Object.fromEntries(request.headers.entries()),
        body,
    };
    console.log(JSON.stringify(response, null, 2));
    return new Response(JSON.stringify(response), { status: 200 });
};

console.log(`HTTP server running. Access it at: http://localhost:${port}/`);
Deno.serve({ port }, handler);
