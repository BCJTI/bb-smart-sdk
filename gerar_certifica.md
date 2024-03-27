### BB Certificado ####
Certificado
openssl pkcs12 -in SmartComex-2023-2024.pfx -clcerts -nokeys -out clientcert.cer
Raiz
openssl pkcs12 -in SmartComex-2023-2024.pfx -cacerts -nokeys -chain -out cacerts.cer
Intermediário
openssl pkcs12 -in SmartComex-2023-2024.pfx -cacerts -nokeys -chain
Key
openssl pkcs12 -in SmartComex-2023-2024.pfx -nocerts -nodes -out clientcert.key