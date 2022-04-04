FROM 127.0.0.1:5000/001

COPY 000.data /003/000-a.data
COPY 001.data /003/000-b.data
COPY 002.data /001/000.data
COPY 000.data /003/000-c.data
