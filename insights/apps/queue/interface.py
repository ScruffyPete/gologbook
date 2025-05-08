import json
from dataclasses import dataclass
from typing import Any, Protocol, Self


class QueueInterface(Protocol):
    async def pop(self) -> Any: ...


@dataclass
class QueueMessage:
    type: str
    payload: dict[str, Any]

    def to_json(self) -> str:
        return json.dumps(self.__dict__)

    @staticmethod
    def from_json(s: str) -> Self:
        return QueueMessage(**json.loads(s))
