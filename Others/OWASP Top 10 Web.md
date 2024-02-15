1. `Broken Access Control`
2. `Cryptographic Failures`
3. `Injection`
4. `Insecure Design`
5. `Security Misconfiguration`
6. `Vulnerable and Outdated Components`
7. `Identification and Authentication Failures`
8. `Software and Data Integrity Failures`
9. `Security Logging and Monitoring Failures`
10. `Server-Side Request Forgery`

------

# 1. Broken Access Control

Broken access control refers to vulnerabilities that allow unauthorized users to access certain resources or perform actions they shouldn't be able to. This can occur due to improper enforcement of restrictions on what authenticated users are allowed to do, or due to flaws in the access control mechanisms themselves.


`Example - IDOR`

## Mitigations

- Implement **role-based access control (RBAC)** to ensure that each user has only the necessary permissions to perform their intended tasks.
- Use **attribute-based access control (ABAC)** to define policies based on specific attributes of the user, object, or environment.
- Regularly **audit access controls** to identify and rectify misconfigurations and inappropriate permissions.
- Enforce **access controls at multiple levels**, including the application, database, and server levels.
- Implement **proper session management and timeout mechanisms** to prevent session hijacking or session fixation attacks.

# 2. Cryptographic Failures

Cryptographic failures occur when cryptographic algorithms, protocols, or keys are not implemented correctly, making them susceptible to attacks. Weak encryption, poor key management, and improper usage of cryptographic libraries are examples of cryptographic failures.

`Example - Getting a database access and able to crack password hashes`
## Mitigations

- Always use **strong and well-established cryptographic algorithms** and protocols.
- Regularly **update cryptographic libraries** and dependencies to avoid known vulnerabilities.
- Implement **proper key management** practices, including secure key generation, storage, rotation, and disposal.
- Use encryption for sensitive data both in transit (e.g., TLS/SSL) and at rest (e.g., disk encryption).
- Perform **regular security assessments and code reviews** to identify and fix cryptographic issues.

# 3. Injection

Injection vulnerabilities occur when untrusted data is improperly handled and gets executed as code. Common types include SQL injection, NoSQL injection, and OS command injection.

`Examples - Command, SQL`

## Mitigations

- Use **parameterized queries** or prepared statements to prevent SQL injection attacks.
- **Sanitize and validate all user-supplied input** to prevent malicious data from being interpreted as code.
- Employ **web application firewalls (WAFs)** to filter and block malicious input.
- Adopt the **principle of least privilege** for database and application users to limit the potential damage of successful injections.

# 4. Insecure Design

Insecure design flaws result from architectural or high-level design choices that make systems vulnerable. These vulnerabilities are typically deeply rooted and can be challenging to fix once the system is developed.

`Example - Password reset vulnerability`

## Mitigations

- Implement secure design principles from the start, such as **defense-in-depth, least privilege, and separation of concerns**.
- Conduct **threat modeling and risk assessments** during the design phase to identify potential vulnerabilities early on.
- Regularly review the design and architecture to ensure they meet security standards and best practices.
- Establish **secure coding guidelines** for developers to follow, with an emphasis on avoiding insecure design patterns.

# 5. Security Misconfiguration

Security misconfiguration occurs when systems, servers, or applications are not properly configured, leaving them exposed to potential attacks. This can include default settings, unnecessary features, and lack of proper access controls.

`Example - Access to interactive console`

## Mitigations

- Follow the **principle of least privilege** and disable unnecessary features and services.
- Keep **software, frameworks, and libraries up to date** to avoid known vulnerabilities.
- Use **strong and unique passwords** for all system components.
- **Regularly audit and review configurations**, using automated tools where possible.
- Implement **proper access controls** and restrict administrative access to only necessary personnel.

# 6. Vulnerable and Outdated Components

Using outdated or vulnerable third-party components (e.g., libraries, plugins, frameworks) can lead to exploitation of known security flaws in these components.

`Example - Publicly available exploits for a specific version of a service`

## Mitigations

- **Regularly update all software and components** to their latest secure versions.
- **Subscribe to security alerts** and mailing lists for third-party components to stay informed about potential vulnerabilities.
- Utilize automated tools to identify and track dependencies, ensuring you can quickly update them when needed.
- Consider using **software composition analysis (SCA)** tools to detect vulnerable components in your codebase.

# 7. Identification and Authentication Failures

Identification and authentication failures involve weaknesses in the way users are identified and how their access is authenticated. This can lead to unauthorized access or impersonation.

`Example - Able to register as an new user and can login as the existing user`

## Mitigations

- Implement multi-factor authentication (MFA) to add an extra layer of security for user logins.
- Use strong password policies, enforce password complexity, and encourage users to choose unique passwords.
- Employ mechanisms like account lockouts and CAPTCHAs to prevent brute-force attacks.
- Regularly review and audit user accounts to remove inactive or unnecessary accounts.

# 8. Software and Data Integrity Failures

Software and data integrity failures occur when data is altered or tampered with, or when malicious code is injected into software or files.

`Example - Websites using third-party libraries stored in some external servers`

`Example - Modifying JWT token cookie using none algorithm attack`
## Mitigations

- Use digital signatures and checksums to verify the integrity of software and updates.
- Employ secure coding practices to prevent code injection vulnerabilities.
- Use secure communication channels for transferring data to prevent tampering during transit (e.g., HTTPS, SFTP).
- Regularly monitor file and system integrity using integrity-checking tools.

# 9. Security Logging and Monitoring Failures

Inadequate security logging and monitoring can lead to delayed or missed detection of security incidents, hindering timely response and investigation.

## Mitigations

- Implement comprehensive logging mechanisms to capture relevant security events and activities.
- Centralize logs and use log management tools to analyze and monitor them effectively.
- Set up real-time alerts to notify administrators of suspicious or anomalous activities.
- Regularly review and analyze logs to identify potential security issues or patterns of attack.

# 10. Server-Side Request Forgery

Server-Side Request Forgery (SSRF) occurs when an attacker manipulates a server into making requests to unintended internal or external systems.

## Mitigations

- Validate and sanitize user-supplied URLs or input that the server processes.
- Use whitelisting to restrict the destinations to which the server can make requests.
- Implement proper access controls and permissions to limit access to sensitive resources.
- Employ network-level filtering to block requests to internal IP addresses from external interfaces.

Addressing these vulnerabilities requires a holistic approach to security, involving secure coding practices, regular security assessments, vulnerability management, and ongoing monitoring and improvements to the system's security posture. Additionally, educating developers and IT personnel about these vulnerabilities and their mitigations is crucial to building a security-conscious culture within an organization.

