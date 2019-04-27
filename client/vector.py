
class vector2:
    def __init__(self, x, y):
        self._x = x
        self._y = y
        pass

    def GetX(self):
        return self._x

    def GetY(self):
        return self._y

    def __repr__(self):
        return 'Point({}, {})'.format(self._x, self._y)

    def __str__(self):
        return '({}, {})'.format(self._x, self._y)

    def __add__(self, other):
        return Point(self._x + other._x, self._y + other._y)

    def __mul__(self, other):
        return Point(self._x * other, self._y * other)

    def __iadd__(self, other):
        self._x += other._x
        self._y += other._y
        return self

    def __sub__(self, other):
        return Point(self._x - other._x, self._y - other._y)

    def __isub__(self, other):
        self._x -= other._x
        self._y -= other._y
        return self

    def __abs__(self):
        return math._ypot(self._x, self._y)

    def __bool__(self):
        return self._x != 0 or self._y != 0

    def __neg__(self):
        return Point(-self._x, -self._y)

    def GetArrInt(self):
        return [int(self._x), int(self._y)]


