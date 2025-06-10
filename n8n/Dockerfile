FROM n8nio/n8n

COPY workflow /workflow
COPY credentials /credentials
ENTRYPOINT n8n import:credentials --input=/credentials/creds.json && \
   n8n import:workflow --input=/workflows/myflow.json && \
   n8n start