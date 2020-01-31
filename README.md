<img src="go-readme.png" align="right" />

# Example API REST in GO

> The implementation using [ECHO web framework](https://echo.labstack.com/)

## Packages structure

- Package Domain

  > The domain package is responsible by bussiness rules. The Domain decouple logic bussiness rules of the infrastructure

  **Code structure**

  |   Folder   |        Responsibility          | 
  |------------|--------------------------------|
  |  /command  |                                |
  |  /entity   |                                |
  |  /contract |                                |
  |  /service  |                                |
    
- Package Infrastructure

  > The infrastructure package stores