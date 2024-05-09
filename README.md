# Warnicke
>Wernicke's area (/ˈvɛərnɪkə/; German: [ˈvɛɐ̯nɪkə]), also called Wernicke's speech area, is one of the two parts of the cerebral cortex that are linked to speech, the other being Broca's area. It is involved in the comprehension of written and spoken language.

`Warnicke` is a small query engine written for learning puposes. It closely follows a similar
spec and architecture found in [datafusion](https://docs.rs/datafusion/latest/datafusion/index.html). 

```
┌─────────────┐    ┌──────────┐      ┌───────────┐     ┌───────────┐
│             │    │          │      │           │     │           │
│  SQL Query  │    │  Parsed  │      │  Logical  │     │ Execution │
│             ├───►│   AST    ├─────►│    Plan   ├────►│   Plan    │
│             │    │          │      │           │     │           │
└─────────────┘    └──────────┘      └───────────┘     └─────┬─────┘
                                                             │      
       Columnar Batches                                      │      
      in Apache Arrow fmt      ┌───────────────┐             │      
      ┌──┐ ┌──┐ ┌──┐ ┌──┐      │               │             │      
      │  │ │  │ │  │ │  │      │     Data      │             │      
      │  │ │  │ │  │ │  │◄─────┤   Source(s)   │◄────────────┘      
      │  │ │  │ │  │ │  │      │               │                    
      └──┘ └──┘ └──┘ └──┘      │               │                    
                               └───────────────┘                    
```

