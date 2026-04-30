# Fractal flame

## Options
- `-w` / `--width` - **int**

  *Width of the final image resolution*

  **Default value**  `1920`
------------------------------------
- `-h` / `--height` - **int**

  *Height of the final image resolution*  
  
  **Default value** `1080`
------------------------------------
- `--seed` - **int**

  *Seed, used in the random generator*  
  
  **Default value** `1`
------------------------------------
- `-i`  /  `--iteration-count` - **int**
  
  *Amount of iterations during generation*
  
  **Default value** `25000`;
------------------------------------
- `--point-iterations` -**int**, 

  *Amount of iterations per random point during generation*

  **Default value** `70`
------------------------------------
- `-o` / `--output-path` - **string**

  *Path to the file, in which the generated image will be written to.*
  *Image is in `PNG` format*

  **Default value** `result.png`
------------------------------------
- `-t` / `--threads` -**int**

  *Amount of threads used during generation*

  **Default value** `1`
------------------------------------
- `-ap` / `--affine-params` - **template string (`aN,bN,cN,dN,eN,fN`)**

  *Affine parameters. Amount of parameters is divisible by 6*

  **Default value** ```-0.6, -0.3, 0.2, -0.15, -0.7, 0.5, 
		-0.3, -0.5, -0.2, 0.5, 0.7, 0.2,
		0.2, -0.35, 0.15, .25, -0.7, 0.8```

  - `a` - **double:** *масштаб/вращение X*
  - `b` - **double:** *сдвиг-смешивание X от Y*
  - `c` - **double:** *сдвиг по X*
  - `d` - **double:** *смешивание Y от X*
  - `e` - **double:** *масштаб/вращение Y*
  - `f` - **double:** *сдвиг по Y*
------------------------------------
- `-f` / `--functions` - **template string (`<function_N>:<function_weight>`)**

  *Non-linear transformations, used during generation*

  **Default value** `swirl:1.0,heart:1.0`
  
  - `<function_N>` - **string** 

    *Function type*
  
  - `<function_weight>` - **double**
  
    *Weight of the functions (its relative probability to be used)*
------------------------------------
- `--config` - **string**

  *Path to the file, containing configuration*

  **Optional argument**

## Config file
Input parameters can be provided in a `.json` file:
```json
{
  "size": {
    "width": 1920, 
    "height": 1080
  },
  "iteration_count": 2500,
  "output_path": "result.png",
  "threads": 4,
  "seed": 5,
  "functions": [
    {
      "name": "swirl",
      "weight": 1.0
    },
    {
      "name": "horseshoe",
      "weight": 0.7
    }
  ],
  "affine_params": {
    "a": 1.0,
    "b": 1.0,
    "c": 1.0,
    "d": 1.0,
    "e": 1.0,
    "f": 1.0
  }
}
```
**Parameter priority:**
* Console args
* JSON-config
* Default values



## Supported non-linear transformations

$$
\begin{aligned} 
& r = \sqrt{x^2 + y^2} &  \\
& \theta = \arctan{\frac{y}{x}} &
\end{aligned}
$$

$$
\begin{aligned} 
& F(x, y) = (x\sin(r^2) - y\cos(r^2) ,\,  x\cos(r^2) + y\sin(r^2)) & \it{Swirl} \\
& F(x, y) = \frac{1}{r} \cdot ((x - y)(x + y) ,\,  2xy)  & \it{Horseshoe} \\
& F(x, y) = \frac{1}{r^2} \cdot (x ,\, y)  & \it{Spherical} \\
& F(x, y) = (\frac{\theta}{\pi} ,\,  r - 1)  & \it{Polar} \\
& F(x, y) = r \cdot(\sin{(r\theta)} ,\, \cos{(r\theta)})  & \it{Heart} \\
& F(x, y) = \frac{\theta}{\pi} \cdot (\sin{(\pi r)} ,\, \cos{(\pi r)})  & \it{Disk} \\
& F(x, y) = (\cos{(\pi x)} \cosh{y} ,\,  \sin{(\pi x)} \sinh{y})  & \it{Cosine} \\
& F(x, y) = (\sin{x} ,\,  \sin{y})  & \it{Sinusoidal} \\
\end{aligned}
$$


## Доп задания

1. 5 и более трансформаций;
2. поддержка логарифмической гамма-коррекции (описано в [статье](https://flam3.com/flame_draves.pdf)):
   * Добавлено поле `gamma_correction` в JSON-конфиг, добавлен параметр `-g`/`--gamma-correction` для CLI:
   * Булевое значение, которое включает или выключает гамма-коррекцию;
   * В случае включения коррекции, яркость пикселей преобразуется по следующей формуле:
   * `color = pow(color, 1.0 / gamma)`, где `gamma` - либо равна `2.2` (значение по умолчанию), либо берется из конфигурации;
   * Добавлено поле `gamma` в JSON-конфиг, добавлен параметр `--gamma` для CLI:
   * Число с плавающей точкой, определяет значение гаммы для коррекции яркости итогового изображения;
   * При наличии параметра в конфигурации, соответствующая логика должна применяться на процесс рендера;
3. поддержка симметрии в генерации пламени (описано в [статье](https://flam3.com/flame_draves.pdf)):
   * Добавлено поле `symmetry_level` в JSON-конфиг, добавлен параметр `-s`/`--symmetry-level` для CLI:
   * Целое число (N >= 1), которое задает количество поворотов вокруг центра;
   * Каждая точка дублируется N раз с поворотом на 360/N градусов;
   * При наличии параметра в конфигурации, соответствующая логика должна применяться на процесс рендера;
