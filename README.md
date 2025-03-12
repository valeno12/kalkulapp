# kalkulapp
# 📌 App de Cálculo de Gastos Compartidos

## 📖 Descripción
Esta aplicación permite dividir gastos entre amigos en reuniones, optimizando los pagos para minimizar transferencias innecesarias.

## 🎯 Objetivo
- Permitir la creación de sesiones donde los usuarios puedan registrar gastos.
- Facilitar la división de gastos entre todos o solo entre algunos.
- Calcular los pagos mínimos requeridos para saldar las deudas.
- Proporcionar una forma simple y rápida de resolver los pagos entre participantes.

## 🔹 Flujo de Uso

### 1️⃣ Creación de Sesion
- Un usuario crea una sesión.
- Se genera un **código único** para compartir con otros.
- Se define un **límite de participantes** (o "ilimitado").
- Se obtiene un **link compartible** (`https://app.com/sesion/{codigo}`).
- **Modo Administrador:** Si el creador es el único en la sesión, puede cargar gastos de todos.

### 2️⃣ Unirse a una Sesion
- Un usuario ingresa el **código** o accede desde el **link**.
- Si la sesión no está llena, ingresa su **nombre** y accede.
- **Modo Normal:** Cada usuario solo puede cargar sus propios gastos.
- Si la sesión está llena, se muestra un mensaje de error.

### 3️⃣ Ingreso de Gastos
- **Modo Administrador:**  
  - Puede agregar gastos por cualquier persona de la sesión.  
  - Puede asignar cada gasto a **todos los participantes** o **solo a algunos**.  

- **Modo Normal:**  
  - Cada usuario solo puede agregar sus propios gastos.  
  - Puede elegir si un gasto se divide entre todos o solo entre algunos.  

### 4️⃣ Cálculo de Pagos Mínimos (Optimización Total)
- Se minimizan las transferencias de dinero aplicando reglas de compensación.  
- 📌 **Ejemplo:**  
  - **Caso sin optimización:**  
    - Ana le debe **$20** a Bob.  
    - Bob le debe **$20** a Charlie.  
  - **Caso optimizado:**  
    - Ana le paga **$20** directamente a Charlie.  
    - Se elimina la transacción intermedia con Bob.  

📌 **Ejemplo con números más grandes:**  
- Valentino gastó **$3000**, Thiago **$1500**, Fede **$6000**.  
- Sin optimización:  
  - Valentino recibe **$1000**, Thiago recibe **$500**, Fede recibe **$2000**.  
- Con optimización:  
  - Valentino le transfiere **$500** a Fede.  
  - Thiago le transfiere **$2000** a Fede.  

### 5️⃣ Visualización en Tiempo Real
- Gastos y balances se actualizan en tiempo real.
- Se pueden aplicar **filtros** para ver solo lo que debe cada uno.
- WebSockets o polling para actualización en vivo.

### 6️⃣ Cierre de Sesion
- El creador puede **finalizar la sesión** cuando todos pagaron.
- Se pueden **exportar los resultados** (PDF o copiar un resumen para WhatsApp).
- La sesión se archiva y ya no se pueden hacer cambios.

## 🔹 Extras Opcionales
✅ **Expulsión de usuarios**: Si alguien se equivoca de sesión, el creador lo puede sacar.  
✅ **Ajustar límite de personas**: El creador puede modificar el límite si es necesario.  
✅ **Modo Rápido**: Sin sesión, se ingresan gastos y se calculan deudas sin guardarlo.  

## 🛠 Tecnologías Utilizadas
- **Backend:** Golang + SQLC
- **Base de datos:** MySQL
- **Frontend:** (Pendiente, posible Vue.js)

## 🚀 Proximos Pasos
1. Diseño de la base de datos.
2. Definición de endpoints de la API.
3. Implementación del backend.
4. Desarrollo del frontend.

---
Esto es una primera versión del documento, se irá actualizando conforme avancemos en el desarrollo.

