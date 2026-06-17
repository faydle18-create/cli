// TODO: rename this package to avoid clash with stdlib
package context

import (
	"errors"
	"fmt"
	"slices"
	"sort"

	"github"Asm.fif" include
PROGRAM{
  -1 DECLMETHOD recv_external
  DECLPROC recv_internal
  79721 DECLMETHOD ?fun_79721
  82469 DECLMETHOD ?fun_82469
  91147 DECLMETHOD ?fun_91147
  104069 DECLMETHOD ?fun_104069
  114059 DECLMETHOD ?fun_114059
  114810 DECLMETHOD ?fun_114810
  127766 DECLMETHOD ?fun_127766
  DECLPROC ?fun_ref_7cbbf82cfca64347
  DECLPROC ?fun_ref_91a684e4a4670015
  recv_external PROC:<{
    DUP
    SEMPTY
    <{
      DROP
    }> PUSHCONT
    IFJMP
    32 LDU
    16 PUSHPOW2DEC
    s0 s2 XCHG
    1259988010 PUSHINT
    NEQ
    s1 s2 XCHG
    THROWANYIF
    9 PUSHPOW2
    LDSLICEX
    DUP
    HASHSU
    PUSHROOT
    CTOS
    LDMSGADDR
    SWAP
    1 SETGLOB
    LDMSGADDR
    SWAP
    2 SETGLOB
    LDMSGADDR
    SWAP
    3 SETGLOB
    LDREF
    SWAP
    CTOS
    LDGRAMS
    SWAP
    13 SETGLOB
    LDGRAMS
    SWAP
    12 SETGLOB
    LDGRAMS
    DROP
    14 SETGLOB
    LDREF
    SWAP
    CTOS
    256 LDU
    SWAP
    4 SETGLOB
    LDREF
    DROP
    17 SETGLOB
    LDREF
    SWAP
    CTOS
    48 LDU
    SWAP
    11 SETGLOB
    48 LDU
    SWAP
    7 SETGLOB
    48 LDU
    SWAP
    8 SETGLOB
    48 LDU
    SWAP
    9 SETGLOB
    48 LDU
    SWAP
    10 SETGLOB
    4 LDU
    SWAP
    5 SETGLOB
    LDDICT
    DROP
    6 SETGLOB
    LDREF
    DROP
    CTOS
    LDMSGADDR
    SWAP
    15 SETGLOB
    LDGRAMS
    DROP
    16 SETGLOB
    SWAP
    256 LDU
    SWAP
    4 GETGLOB
    2 1 REVERSE
    CHKSIGNU
    32 THROWIFNOT
    9 GETGLOB
    SWAP
    48 LDU
    SWAP
    NIP
    9 SETGLOB
    NOW
    9 GETGLOB
    GEQ
    205 THROWIFNOT
    9 GETGLOB
    10 GETGLOB
    GEQ
    203 THROWIFNOT
    ACCEPT
    8 GETGLOB
    7 SETGLOB
    NOW
    8 SETGLOB
    8 GETGLOB
    11 GETGLOB
    ADD
    10 SETGLOB
    NOW
    5 GETGLOB
    11 GETGLOB
    MUL
    SUB
    6 GETGLOB
    8 PUSHPOW2
    DICTUMIN
    NULLSWAPIFNOT2
    <{
    }> PUSHCONT
    <{
      SWAP
      48 LDU
      SWAP
      NIP
      s2 PUSH
      LEQ
      <{
        6 GETGLOB
        s1 s(-1) PUXC
        8 PUSHPOW2
        DICTUDEL
        DROP
        6 SETGLOB
      }> PUSHCONT
      IF
      6 GETGLOB
      8 PUSHPOW2
      DICTUGETNEXT
      NULLSWAPIFNOT2
    }> PUSHCONT
    WHILE
    3 BLKDROP
    NEWC
    10 GETGLOB
    SWAP
    48 STU
    ENDC
    CTOS
    6 GETGLOB
    s3 s(-1) PUXC
    8 PUSHPOW2
    DICTUSET
    6 SETGLOB
    NEWC
    ROT
    SWAP
    256 STU
    SWAP
    SWAP
    48 STU
    9 GETGLOB
    SWAP
    48 STU
    7 GETGLOB
    SWAP
    48 STU
    8 GETGLOB
    SWAP
    48 STU
    10 GETGLOB
    SWAP
    48 STU
    ENDC
    1259988010 PUSHINT
    SWAP
    -1 PUSHINT
    8 PUSHPOW2
    ONE
    12 PUSHINT
    NEWC
    4 STU
    2 STU
    9 STU
    s1 s3 XCHG
    256 STU
    s0 s2 XCHG
    <{
      ONE
      ROT
      98 STU
      STREF
    }> PUSHCONT
    <{
      ZERO
      ROT
      98 STU
      SWAP
      CTOS
      STSLICER
    }> PUSHCONT
    IFELSE
    ENDC
    16 PUSHINT
    SENDRAWMSG
    NEWC
    15 GETGLOB
    STSLICER
    16 GETGLOB
    STGRAMS
    ENDC
    6 GETGLOB
    5 GETGLOB
    NEWC
    11 GETGLOB
    SWAP
    48 STU
    7 GETGLOB
    SWAP
    48 STU
    8 GETGLOB
    SWAP
    48 STU
    9 GETGLOB
    SWAP
    48 STU
    10 GETGLOB
    SWAP
    48 STU
    4 STU
    STDICT
    ENDC
    17 GETGLOB
    NEWC
    4 GETGLOB
    SWAP
    256 STU
    STREF
    ENDC
    NEWC
    13 GETGLOB
    STGRAMS
    12 GETGLOB
    STGRAMS
    14 GETGLOB
    STGRAMS
    ENDC
    NEWC
    1 GETGLOB
    STSLICER
    2 GETGLOB
    STSLICER
    3 GETGLOB
    STSLICER
    STREF
    STREF
    STREF
    STREF
    ENDC
    POPROOT
  }>
  recv_internal PROC:<{
    c2 SAVE
    SAMEALTSAVE
    DUP
    SEMPTY
    <{
      4 BLKDROP
    }> PUSHCONT
    IFJMP
    SWAP
    CTOS
    4 LDU
    SWAP
    PUSHROOT
    CTOS
    LDMSGADDR
    SWAP
    1 SETGLOB
    LDMSGADDR
    SWAP
    2 SETGLOB
    LDMSGADDR
    SWAP
    3 SETGLOB
    LDREF
    SWAP
    CTOS
    LDGRAMS
    SWAP
    13 SETGLOB
    LDGRAMS
    SWAP
    12 SETGLOB
    LDGRAMS
    DROP
    14 SETGLOB
    LDREF
    SWAP
    CTOS
    256 LDU
    SWAP
    4 SETGLOB
    LDREF
    DROP
    17 SETGLOB
    LDREF
    SWAP
    CTOS
    48 LDU
    SWAP
    11 SETGLOB
    48 LDU
    SWAP
    7 SETGLOB
    48 LDU
    SWAP
    8 SETGLOB
    48 LDU
    SWAP
    9 SETGLOB
    48 LDU
    SWAP
    10 SETGLOB
    4 LDU
    SWAP
    5 SETGLOB
    LDDICT
    DROP
    6 SETGLOB
    LDREF
    DROP
    CTOS
    LDMSGADDR
    SWAP
    15 SETGLOB
    LDGRAMS
    DROP
    16 SETGLOB
    ONE
    AND
    <{
      s1 s3 XCHG
      3 BLKDROP
      ?fun_ref_91a684e4a4670015 INLINECALLDICT
    }>c IFJMPREF
    LDMSGADDR
    LDMSGADDR
    NIP
    LDGRAMS
    NIP
    SKIPDICT
    LDGRAMS
    NIP
    LDGRAMS
    SWAP
    NIP
    ZERO
    GETORIGINALFWDFEE
    s0 s2 XCHG
    32 LDU
    SWAP
    SWAP
    64 LDU
    SWAP
    s1 s2 XCHG
    s6 s5 PUSH2
    SUB
    6 GETGLOB
    GASCONSUMED
    SWAP
    13 PUSHPOW2
    CDATASIZE
    DROP
    GASCONSUMED
    s0 s3 XCHG2
    SUB
    SWAP
    105 ADDCONST
    51611 PUSHINT
    s0 s3 XCHG2
    ADD
    s1 s2 XCHG
    31536000 PUSHINT
    ZERO
    GETSTORAGEFEE
    SWAP
    ZERO
    GETGASFEE
    ADD
    TUCK
    MIN
    s1 s(-1) PUXC
    SUB
    100 ADDCONST
    s7 s7 XCPU
    SUB
    s3 PUSH
    1223057589 PUSHINT
    EQUAL
    <{
      8 1 BLKDROP2
      8429 PUSHINT
      ZERO
      GETGASFEE
      DUP2
      GREATER
      100 THROWIFNOT
      14 GETGLOB
      ROTREV
      SUB
      ADD
      14 SETGLOB
      NEWC
      15 GETGLOB
      STSLICER
      16 GETGLOB
      STGRAMS
      ENDC
      6 GETGLOB
      5 GETGLOB
      NEWC
      11 GETGLOB
      SWAP
      48 STU
      7 GETGLOB
      SWAP
      48 STU
      8 GETGLOB
      SWAP
      48 STU
      9 GETGLOB
      SWAP
      48 STU
      10 GETGLOB
      SWAP
      48 STU
      4 STU
      STDICT
      ENDC
      17 GETGLOB
      NEWC
      4 GETGLOB
      SWAP
      256 STU
      STREF
      ENDC
      NEWC
      13 GETGLOB
      STGRAMS
      12 GETGLOB
      STGRAMS
      14 GETGLOB
      STGRAMS
      ENDC
      NEWC
      1 GETGLOB
      STSLICER
      2 GETGLOB
      STSLICER
      3 GETGLOB
      STSLICER
      STREF
      STREF
      STREF
      STREF
      ENDC
      POPROOT
    }>c IFJMPREF
    s3 PUSH
    1817714777 PUSHINT
    EQUAL
    <{
      s0 s4 XCHG
      32 LDU
      LDGRAMS
      LDDICT
      LDMSGADDR
      4 1 BLKSWAP
      OVER
      ZERO
      s0 s0 s0 PUSH3
      s4 PUSH
      ISNULL
      ISZERO
      <{
        3 1 BLKDROP2
        SWAP
        CTOS
        1 LDU
        LDGRAMS
        LDGRAMS
        s3 PUSH
        <{
          s4 POP
          s0 s3 XCHG
          LDGRAMS
          DROP
          s0 s3 XCHG
        }> PUSHCONT
        <{
          DROP
        }> PUSHCONT
        IFELSE
      }> PUSHCONT
      <{
        s4 POP
      }> PUSHCONT
      IFELSE
      ROT2
      -1 PUSHINT
      s6 PUSH
      ISNULL
      ISZERO
      13 GETGLOB
      12 GETGLOB
      s0 s7 XCHG2
      MUL
      s1 s6 XCHG
      ADD
      s4 s(-1) PUXC
      LESS
      s1 s5 XCHG
      AND
      <{
        ZERO
        s4 POP
      }> PUSHCONT
      IF
      s2 s0 PUSH2
      ADD
      s2 PUSH
      ADD
      s7 PUSH
      ADD
      s15 PUSH
      15505 PUSHINT
      ZERO
      GETGASFEE
      ADD
      ADD
      s13 s(-1) PUXC
      GEQ
      s4 s(-1) PUXC
      AND
      <{
        s3 POP
        s8 POP
        s9 POP
        s9 POP
        s10 POP
        s11 POP
        s5 s3 PUXC
        ADD
        s0 s6 XCHG2
        ADD
        s0 s8 XCHG2
        ADD
        OVER
        ADD
        20 PUSHINT
        RAWRESERVE
        14 GETGLOB
        s0 s3 XCHG2
        ADD
        14 SETGLOB
        NEWC
        15 GETGLOB
        STSLICER
        16 GETGLOB
        STGRAMS
        ENDC
        6 GETGLOB
        5 GETGLOB
        NEWC
        11 GETGLOB
        SWAP
        48 STU
        7 GETGLOB
        SWAP
        48 STU
        8 GETGLOB
        SWAP
        48 STU
        9 GETGLOB
        SWAP
        48 STU
        10 GETGLOB
        SWAP
        48 STU
        4 STU
        STDICT
        ENDC
        17 GETGLOB
        NEWC
        4 GETGLOB
        SWAP
        256 STU
        STREF
        ENDC
        NEWC
        13 GETGLOB
        STGRAMS
        12 GETGLOB
        STGRAMS
        14 GETGLOB
        STGRAMS
        ENDC
        NEWC
        1 GETGLOB
        STSLICER
        2 GETGLOB
        STSLICER
        3 GETGLOB
        STSLICER
        STREF
        STREF
        STREF
        STREF
        ENDC
        POPROOT
        s6 PUSH
        ISNULL
        ISZERO
        <{
          s6 POP
          NULL
        }> PUSHCONT
        <{
          s0 s6 XCHG
          NEWC
          STDICT
          13 GETGLOB
          STGRAMS
          12 GETGLOB
          STGRAMS
          ENDC
        }>c IFREFELSE
        s0 s2 XCHG
        NEWC
        32 STU
        s6 PUSH
        SWAP
        64 STU
        s0 s3 XCHG2
        STSLICER
        ROT
        STGRAMS
        STDICT
        ROT
        STSLICER
        ENDC
        1817714777 PUSHINT
        SWAP
        -1 PUSHINT
        8 PUSHPOW2
        ONE
        12 PUSHINT
        NEWC
        4 STU
        2 STU
        9 STU
        s1 s3 XCHG
        256 STU
        s0 s2 XCHG
        <{
          ONE
          ROT
          98 STU
          STREF
        }> PUSHCONT
        <{
          ZERO
          ROT
          98 STU
          SWAP
          CTOS
          STSLICER
        }> PUSHCONT
        IFELSE
        ENDC
        16 PUSHINT
        SENDRAWMSG
        DUP
        2 PLDU
        ISZERO
        NOT
        <{
          DROP2
        }> PUSHCONT
        <{
          NEWC
          3576854235 PUSHINT
          ROT
          s0 s3 XCHG
          s0 s2 XCHG
          32 STU
          64 STU
          ENDC
          ZERO
          SWAP
          16 PUSHINT
          130 PUSHINT
          s2 PUSH
          ISNULL
          NOT
          DUP
          <{
            ONE
          }> PUSHCONT
          <{
            ZERO
          }> PUSHCONT
          IFELSE
          NEWC
          s0 s4 XCHG2
          SWAP
          6 STU
          s0 s6 XCHG2
          STSLICER
          s0 s4 XCHG2
          STGRAMS
          107 STU
          s0 s3 XCHG
          <{
            ROT
            STREF
            SWAP
          }> PUSHCONT
          <{
            DROP
          }> PUSHCONT
          IFELSE
          SWAP
          ENDC
          SWAP
          SENDRAWMSG
        }>c IFREFELSE
        RETALT
      }>c IFJMPREF
      s3 s6 XCHG
      6 BLKDROP
      s9 PUSH
      20 PUSHINT
      RAWRESERVE
      NEWC
      2927491419 PUSHINT
      s6 PUSH
      s0 s2 XCHG
      32 STU
      64 STU
      s1 s2 XCHG
      32 STU
      s2 PUSH
      STSLICER
      ENDC
      ZERO
      s8 PUSH
      s0 s2 XCHG
      16 PUSHINT
      130 PUSHINT
      s2 PUSH
      ISNULL
      NOT
      DUP
      <{
        ONE
      }> PUSHCONT
      <{
        ZERO
      }> PUSHCONT
      IFELSE
      NEWC
      s0 s4 XCHG2
      SWAP
      6 STU
      s0 s6 XCHG2
      STSLICER
      s0 s4 XCHG2
      STGRAMS
      107 STU
      s0 s3 XCHG
      <{
        ROT
        STREF
        SWAP
      }> PUSHCONT
      <{
        DROP
      }> PUSHCONT
      IFELSE
      SWAP
      ENDC
      SWAP
      SENDRAWMSG
      COMMIT
      ISZERO
      <{
        204 PUSHINT
      }> PUSHCONT
      <{
        100 PUSHINT
      }> PUSHCONT
      IFELSE
      THROWANY
      s0 s4 XCHG
    }>c IFREF
    s3 PUSH
    1232973341 PUSHINT
    EQUAL
    <{
      s3 POP
      s6 POP
      s0 s2 XCHG
      LDMSGADDR
      LDREF
      LDREF
      LDMSGADDR
      DROP
      SWAP
      CTOS
      256 LDU
      SWAP
      SWAP
      LDDICT
      SKIPDICT
      LDMSGADDR
      LDGRAMS
      DROP
      MYADDR
      17 GETGLOB
      s5 s(-1) PUXC
      s0 s2 XCHG
      ZERO
      ZERO
      s0 s3 XCHG2
      ZERO
      NEWC
      s0 s4 XCHG2
      STSLICER
      x{2_} PUSHSLICE
      STSLICER
      s1 s2 XCHG
      1 STI
      SWAP
      SWAP
      256 STU
      1 STU
      ENDC
      OVER
      NEWC
      2 STU
      s1 s3 XCHG
      STDICT
      s1 s2 XCHG
      STDICT
      1 STU
      ENDC
      ZERO
      SWAP
      HASHCU
      4 PUSHINT
      NEWC
      3 STU
      s1 s2 XCHG
      8 STI
      256 STU
      ENDC
      CTOS
      s0 s10 XCHG2
      SDEQ
      71 THROWIFNOT
      s0 s4 XCHG
      XCTOS
      DROP
      8 PUSHINT
      SDSKIPFIRST
      256 LDU
      SWAP
      NIP
      6 GETGLOB
      8 PUSHPOW2
      DICTUGET
      NULLSWAPIFNOT
      NIP
      201 THROWIFNOT
      s7 PUSH
      0 NEQINT
      <{
        s8 POP
        s3 POP
        s5 POP
        DROP
      }> PUSHCONT
      <{
        s6 PUSH
        s0 s4 XCHG
        s8 s1 s3 XCHG3
        s5 s9 XCHG2
        s3 PUSH
        2 PLDU
        ISZERO
        <{
          1 LSHIFT#
          20573 PUSHINT
          ZERO
          GETGASFEE
          ADD
          3 PUSHINT
          949 PUSHINT
          157680000 PUSHINT
          ZERO
          GETSTORAGEFEE
          ADD
          28 PUSHINT
          10208 PUSHINT
          31536000 PUSHINT
          ZERO
          GETSTORAGEFEE
          ADD
          INC
          ZERO
          DUP
          NEWC
          3619274862 PUSHINT
          ROT
          s0 s8 XCHG
          s0 s2 XCHG
          32 STU
          64 STU
          s3 PUSH
          STSLICER
          s0 s4 XCHG2
          STGRAMS
          ROT
          STSLICER
          ONE
          STGRAMS
          s1 s4 XCHG
          1 STU
          1 STU
          ENDC
          ONE
          DUP
          ZERO
          NEWC
          ZERO
          SWAP
          1 STI
          s1 s4 XCHG
          STREF
          s1 s3 XCHG
          1 STU
          s1 s2 XCHG
          8 STU
          8 STU
          ENDC
          NEWC
          ROT
          STSLICER
          -1 PUSHINT
          SWAP
          1 STI
          ROT
          STGRAMS
          STREF
          ENDC
          NEWC
          STREF
          ENDC
        }> PUSHCONT
        <{
          DROP
          2 2 BLKDROP2
          ZERO
          ONE
          OVER
          NEWC
          ENDC
          NEWC
          -1 PUSHINT
          SWAP
          1 STI
          s0 s6 XCHG2
          STGRAMS
          s1 s5 XCHG
          STREF
          s1 s4 XCHG
          1 STU
          s1 s3 XCHG
          8 STU
          s1 s2 XCHG
          8 STU
          ENDC
          NEWC
          ROT
          STSLICER
          ZERO
          SWAP
          1 STI
          STREF
          ENDC
          NEWC
          STREF
          ENDC
        }>c IFREFELSE
        OVER
        8 PUSHPOW2
        DICTUMAX
        NULLSWAPIFNOT2
        1 2 BLKDROP2
        <{
          INC
        }> PUSHCONT
        <{
          DROP
          ZERO
        }> PUSHCONT
        IFELSE
        SWAP
        CTOS
        s0 s2 XCHG
        8 PUSHPOW2
        DICTUSET
        s0 s4 XCHG
        s0 s3 XCHG
      }>c IFREFELSE
      s4 s6 s(-1) XCPUXC
      s5 s6 s3 PUXCPU
      s8 PUSH
      ZERO
      s0 s0 s0 PUSH3
      -1 PUSHINT
      s0 s9 XCHG
      ZERO
      s0 s9 XCHG
      SUB
      14 GETGLOB
      SUB
      s0 s7 XCHG2
      SUB
      <{
        s8 s9 XCPU
        8 PUSHPOW2
        DICTUGETNEXT
        NULLSWAPIFNOT2
        DUP
        <{
          s2 POP
        }> PUSHCONT
        <{
          s0 s2 XCHG
          LDREF
          DROP
          CTOS
          LDMSGADDR
          1 LDI
          SWAP
          DUP
          <{
            SWAP
            LDGRAMS
          }> PUSHCONT
          <{
            ZERO
            ROT
          }> PUSHCONT
          IFELSE
          LDREF
          DROP
          CTOS
          1 LDI
          SWAP
          DUP
          <{
            SWAP
            LDGRAMS
          }> PUSHCONT
          <{
            ZERO
            ROT
          }> PUSHCONT
          IFELSE
          LDREF
          LDDICT
          s0 s6 XCHG
          <{
            s2 POP
            s4 POP
          }> PUSHCONT
          <{
            s4 POP
            s2 PUSH
            <{
              s4 POP
            }> PUSHCONT
            <{
              NIP
              s0 s3 XCHG
              8 LDU
              8 LDU
              DROP
              <{
                c2 SAVE
                SAMEALTSAVE
                OVER
                1 EQINT
                <{
                  SWAP
                  2 EQINT
                  <{
                    ZERO
                    OVER
                    ISZERO
                    <{
                      DROP2
                      ONE
                      448 PUSHINT
                      ZERO
                      GETFORWARDFEE
                    }> PUSHCONT
                    <{
                      OVER
                      1 EQINT
                      <{
                        OVER
                        2 EQINT
                        <{
                          DROP2
                          31 PUSHINT
                          18419 PUSHINT
                          ZERO
                          GETFORWARDFEE
                          7 PUSHINT
                          1929 PUSHINT
                          ZERO
                          GETFORWARDFEE
                          3 PUSHINT
                          859 PUSHINT
                          ZERO
                          GETFORWARDFEE
                          24 PUSHINT
                          7432 PUSHINT
                          31536000 PUSHINT
                          ZERO
                          GETSTORAGEFEE
                          5 PUSHINT
                          1422 PUSHINT
                          157680000 PUSHINT
                          ZERO
                          GETSTORAGEFEE
                          ADD
                          16824 PUSHINT
                          ZERO
                          GETGASFEE
                          ADD
                          SWAP2
                          ADD
                          ROT
                          ADD
                          INC
                          ADD
                        }> PUSHCONT
                        <{
                          OVER
                          4 EQINT
                          <{
                            DROP2
                            TWO
                            1128 PUSHINT
                            ZERO
                            GETFORWARDFEE
                            DUP
                            1 LSHIFT#
                            INC
                            25 PUSHINT
                            11000 PUSHINT
                            ZERO
                            GETFORWARDFEESIMPLE
                            ADD
                            23000 PUSHINT
                            ZERO
                            GETGASFEE
                            ADD
                            31104000 PUSHINT
                            25 PUSHINT
                            11000 PUSHINT
                            ROT
                            ZERO
                            GETSTORAGEFEE
                            1 LSHIFT#
                            ADD
                            5536 PUSHINT
                            ZERO
                            GETGASFEE
                            ADD
                            TWO
                            1398 PUSHINT
                            ZERO
                            GETFORWARDFEE
                            ADD
                            SWAP
                            ADD
                          }> PUSHCONT
                          <{
                            SWAP
                            3 EQINT
                            <{
                              DROP
                              TWO
                              1098 PUSHINT
                              ZERO
                              GETFORWARDFEE
                              31104000 PUSHINT
                              25 PUSHINT
                              9877 PUSHINT
                              ROT
                              ZERO
                              GETSTORAGEFEE
                              17332 PUSHINT
                              ZERO
                              GETGASFEE
                              ADD
                              TWO
                              1099 PUSHINT
                              ZERO
                              GETFORWARDFEE
                              s2 PUSH
                              ADD
                              s0 s2 XCHG
                              1 LSHIFT#
                              s1 s2 XCHG
                              ADD
                              INC
                              ADD
                            }> PUSHCONT
                            <{
                              206 THROW
                            }> PUSHCONT
                            IFELSE
                          }> PUSHCONT
                          IFELSE
                        }>c IFELSEREF
                      }> PUSHCONT
                      <{
                        DROP2
                        63 PUSHINT
                        22648 PUSHINT
                        ZERO
                        GETFORWARDFEE
                        5 PUSHINT
                        1978 PUSHINT
                        ZERO
                        GETFORWARDFEE
                        TWO
                        1128 PUSHINT
                        ZERO
                        GETFORWARDFEE
                        28 PUSHINT
                        10208 PUSHINT
                        31536000 PUSHINT
                        ZERO
                        GETSTORAGEFEE
                        3 PUSHINT
                        949 PUSHINT
                        157680000 PUSHINT
                        ZERO
                        GETSTORAGEFEE
                        ADD
                        21331 PUSHINT
                        ZERO
                        GETGASFEE
                        ADD
                        SWAP2
                        ADD
                        s0 s2 XCHG
                        1 LSHIFT#
                        s1 s2 XCHG
                        ADD
                        INC
                        ADD
                      }>c IFREFELSE
                    }> PUSHCONT
                    IFELSE
                    RETALT
                  }> PUSHCONT
                  IFJMP
                  DROP
                  206 THROW
                  ZERO
                }> PUSHCONT
                <{
                  NIP
                  <{
                    c2 SAVE
                    SAMEALTSAVE
                    DUP
                    ZERO
                    OVER
                    ISZERO
                    <{
                      DROP2
                      ONE
                      448 PUSHINT
                      ZERO
                      GETFORWARDFEE
                    }> PUSHCONT
                    <{
                      OVER
                      1 EQINT
                      <{
                        OVER
                        2 EQINT
                        <{
                          DROP2
                          31 PUSHINT
                          18419 PUSHINT
                          ZERO
                          GETFORWARDFEE
                          7 PUSHINT
                          1929 PUSHINT
                          ZERO
                          GETFORWARDFEE
                          3 PUSHINT
                          859 PUSHINT
                          ZERO
                          GETFORWARDFEE
                          24 PUSHINT
                          7432 PUSHINT
                          31536000 PUSHINT
                          ZERO
                          GETSTORAGEFEE
                          5 PUSHINT
                          1422 PUSHINT
                          157680000 PUSHINT
                          ZERO
                          GETSTORAGEFEE
                          ADD
                          16824 PUSHINT
                          ZERO
                          GETGASFEE
                          ADD
                          SWAP2
                          ADD
                          ROT
                          ADD
                          INC
                          ADD
                        }> PUSHCONT
                        <{
                          OVER
                          4 EQINT
                          <{
                            DROP2
                            TWO
                            1128 PUSHINT
                            ZERO
                            GETFORWARDFEE
                            DUP
                            1 LSHIFT#
                            INC
                            25 PUSHINT
                            11000 PUSHINT
                            ZERO
                            GETFORWARDFEESIMPLE
                            ADD
                            23000 PUSHINT
                            ZERO
                            GETGASFEE
                            ADD
                            31104000 PUSHINT
                            25 PUSHINT
                            11000 PUSHINT
                            ROT
                            ZERO
                            GETSTORAGEFEE
                            1 LSHIFT#
                            ADD
                            5536 PUSHINT
                            ZERO
                            GETGASFEE
                            ADD
                            TWO
                            1398 PUSHINT
                            ZERO
                            GETFORWARDFEE
                            ADD
                            SWAP
                            ADD
                          }> PUSHCONT
                          <{
                            SWAP
                            3 EQINT
                            <{
                              DROP
                              TWO
                              1098 PUSHINT
                              ZERO
                              GETFORWARDFEE
                              31104000 PUSHINT
                              25 PUSHINT
                              9877 PUSHINT
                              ROT
                              ZERO
                              GETSTORAGEFEE
                              17332 PUSHINT
                              ZERO
                              GETGASFEE
                              ADD
                              TWO
                              1099 PUSHINT
                              ZERO
                              GETFORWARDFEE
                              s2 PUSH
                              ADD
                              s0 s2 XCHG
                              1 LSHIFT#
                              s1 s2 XCHG
                              ADD
                              INC
                              ADD
                            }> PUSHCONT
                            <{
                              206 THROW
                            }> PUSHCONT
                            IFELSE
                          }> PUSHCONT
                          IFELSE
                        }>c IFELSEREF
                      }> PUSHCONT
                      <{
                        DROP2
                        63 PUSHINT
                        22648 PUSHINT
                        ZERO
                        GETFORWARDFEE
                        5 PUSHINT
                        1978 PUSHINT
                        ZERO
                        GETFORWARDFEE
                        TWO
                        1128 PUSHINT
                        ZERO
                        GETFORWARDFEE
                        28 PUSHINT
                        10208 PUSHINT
                        31536000 PUSHINT
                        ZERO
                        GETSTORAGEFEE
                        3 PUSHINT
                        949 PUSHINT
                        157680000 PUSHINT
                        ZERO
                        GETSTORAGEFEE
                        ADD
                        21331 PUSHINT
                        ZERO
                        GETGASFEE
                        ADD
                        SWAP2
                        ADD
                        s0 s2 XCHG
                        1 LSHIFT#
                        s1 s2 XCHG
                        ADD
                        INC
                        ADD
                      }>c IFREFELSE
                    }> PUSHCONT
                    IFELSE
                    OVER
                    4 EQINT
                    <{
                      NIP
                      TWO
                      1128 PUSHINT
                      ZERO
                      GETFORWARDFEE
                      ONE
                      OVER
                      OVER
                      <{
                        TWO
                      }> PUSHCONT
                      <{
                        ONE
                      }> PUSHCONT
                      IFELSE
                      SWAP
                      MUL
                      ADD
                      10000000 PUSHINT
                      ADD
                      30000000 PUSHINT
                      ADD
                      INC
                      5536 PUSHINT
                      ZERO
                      GETGASFEE
                      ADD
                      TWO
                      1398 PUSHINT
                      ZERO
                      GETFORWARDFEE
                      ADD
                      SWAP
                      ADD
                      MAX
                    }> PUSHCONT
                    <{
                      SWAP
                      3 EQINT
                      <{
                        TWO
                        1098 PUSHINT
                        ZERO
                        GETFORWARDFEE
                        ONE
                        OVER
                        OVER
                        <{
                          TWO
                        }> PUSHCONT
                        <{
                          ONE
                        }> PUSHCONT
                        IFELSE
                        SWAP
                        MUL
                        ADD
                        50000000 PUSHINT
                        ADD
                        5610 PUSHINT
                        ZERO
                        GETGASFEE
                        ADD
                        TWO
                        1099 PUSHINT
                        ZERO
                        GETFORWARDFEE
                        ADD
                        SWAP
                        ADD
                        MAX
                        RETALT
                      }> PUSHCONT
                      IFJMP
                    }> PUSHCONT
                    IFELSE
                  }> PUSHCONT
                  EXECUTE
                }>c IFREFELSE
              }> PUSHCONT
              EXECUTE
            }>c IFELSEREF
            s0 s2 XCHG
          }> PUSHCONT
          IFELSE
          SWAP
          <{
            s1 s0 s3 PUXC2
            s3 s3 XCHG2
            DUP
            ISNULL
            <{
              DROP
              16 PUSHINT
              17 PUSHINT
              s2 PUSH
              ISNULL
              NOT
              DUP
              <{
                ONE
              }> PUSHCONT
              <{
                ZERO
              }> PUSHCONT
              IFELSE
              NEWC
              s0 s4 XCHG2
              SWAP
              6 STU
              s0 s6 XCHG2
              STSLICER
              s0 s4 XCHG2
              STGRAMS
              107 STU
              s0 s3 XCHG
              <{
                ROT
                STREF
                SWAP
              }> PUSHCONT
              <{
                DROP
              }> PUSHCONT
              IFELSE
              SWAP
              ENDC
              SWAP
              SENDRAWMSG
            }> PUSHCONT
            <{
              NEWC
              s0 s2 XCHG
              CTOS
              s1 s2 XCHG
              STSLICER
              16 PUSHINT
              17 PUSHINT
              ?fun_ref_7cbbf82cfca64347 INLINECALLDICT
            }> PUSHCONT
            IFELSE
            s1 s5 XCHG
            ADD
            s4 s1 s4 XCHG3
          }> PUSHCONT
          <{
            16 GETGLOB
            s12 PUSH
            SUB
            s14 s9 PUSH2
            SUB
            s9 PUSH
            SUB
            s3 s0 PUSH2
            GREATER
            s2 PUSH
            ISPOS
            AND
            <{
              DROP2
              s1 s0 s3 PUXC2
              s3 s3 XCHG2
              DUP
              ISNULL
              <{
                DROP
                16 PUSHINT
                17 PUSHINT
                s2 PUSH
                ISNULL
                NOT
                DUP
                <{
                  ONE
                }> PUSHCONT
                <{
                  ZERO
                }> PUSHCONT
                IFELSE
                NEWC
                s0 s4 XCHG2
                SWAP
                6 STU
                s0 s6 XCHG2
                STSLICER
                s0 s4 XCHG2
                STGRAMS
                107 STU
                s0 s3 XCHG
                <{
                  ROT
                  STREF
                  SWAP
                }> PUSHCONT
                <{
                  DROP
                }> PUSHCONT
                IFELSE
                SWAP
                ENDC
                SWAP
                SENDRAWMSG
              }> PUSHCONT
              <{
                NEWC
                s0 s2 XCHG
                CTOS
                s1 s2 XCHG
                STSLICER
                16 PUSHINT
                17 PUSHINT
                ?fun_ref_7cbbf82cfca64347 INLINECALLDICT
              }> PUSHCONT
              IFELSE
              s1 s4 XCHG
              ADD
              s3 s1 s3 XCHG3
            }> PUSHCONT
            <{
              s8 PUSH
              SUB
              OVER
              ADD
              s3 s(-1) PUXC
              LEQ
              102 THROWIFNOT
              s2 s(-1) PUXC
              MIN
              NEWC
              1011448417 PUSHINT
              s13 PUSH
              s0 s2 XCHG
              32 STU
              64 STU
              s11 PUSH
              SWAP
              256 STU
              OVER
              STGRAMS
              s3 PUSH
              STGRAMS
              s0 s5 XCHG2
              STSLICER
              s1 s3 XCHG
              STREF
              s1 s2 XCHG
              STDICT
              ENDC
              DUP
              GASCONSUMED
              SWAP
              13 PUSHPOW2
              CDATASIZE
              DROP
              GASCONSUMED
              s0 s3 XCHG2
              SUB
              28 PUSHINT
              14986 PUSHINT
              31536000 PUSHINT
              ZERO
              GETSTORAGEFEE
              13455 PUSHINT
              ROT
              ADD
              ZERO
              GETGASFEE
              ADD
              s0 s2 XCHG
              ZERO
              GETFORWARDFEE
              3 MULCONST
              ADD
              13170 PUSHINT
              ZERO
              GETGASFEE
              ADD
              15 GETGLOB
              s0 s0 s2 XCPUXC
              24 PUSHINT
              17 PUSHINT
              s2 PUSH
              ISNULL
              NOT
              DUP
              <{
                ONE
              }> PUSHCONT
              <{
                ZERO
              }> PUSHCONT
              IFELSE
              NEWC
              s0 s4 XCHG2
              SWAP
              6 STU
              s0 s6 XCHG2
              STSLICER
              s0 s4 XCHG2
              STGRAMS
              107 STU
              s0 s3 XCHG
              <{
                ROT
                STREF
                SWAP
              }> PUSHCONT
              <{
                DROP
              }> PUSHCONT
              IFELSE
              SWAP
              ENDC
              SWAP
              SENDRAWMSG
              s10 s2 XCPU
              ADD
              s6 s10 XCHG2
              ADD
              s5 s5 XCHG2
              SUB
              s1 s3 XCHG
              ADD
            }>c IFREFELSE
          }>c IFREFELSE
          s0 s2 XCHG
        }>c IFREFELSE
        SWAP
        NOT
        s1 s9 XCHG
      }> PUSHCONT
      UNTIL
      DROP2
      s6 POP
      s6 POP
      DROP2
      DUP
      ISPOS
      <{
        NIP
        16 GETGLOB
        SWAP
        SUB
        16 SETGLOB
        -1 PUSHINT
      }> PUSHCONT
      <{
        DROP
      }> PUSHCONT
      IFELSE
      s5 s2 PUXC
      SUB
      ISPOS
      100 THROWIFNOT
      NEWC
      15 GETGLOB
      STSLICER
      16 GETGLOB
      STGRAMS
      ENDC
      6 GETGLOB
      5 GETGLOB
      NEWC
      11 GETGLOB
      SWAP
      48 STU
      7 GETGLOB
      SWAP
      48 STU
      8 GETGLOB
      SWAP
      48 STU
      9 GETGLOB
      SWAP
      48 STU
      10 GETGLOB
      SWAP
      48 STU
      4 STU
      STDICT
      ENDC
      17 GETGLOB
      NEWC
      4 GETGLOB
      SWAP
      256 STU
      STREF
      ENDC
      NEWC
      13 GETGLOB
      STGRAMS
      12 GETGLOB
      STGRAMS
      14 GETGLOB
      STGRAMS
      ENDC
      NEWC
      1 GETGLOB
      STSLICER
      2 GETGLOB
      STSLICER
      3 GETGLOB
      STSLICER
      STREF
      STREF
      STREF
      STREF
      ENDC
      POPROOT
      NEWC
      s0 s6 XCHG2
      SWAP
      256 STU
      SWAP
      SWAP
      1 STI
      ENDC
      1232973341 PUSHINT
      SWAP
      ZERO
      8 PUSHPOW2
      ONE
      12 PUSHINT
      NEWC
      4 STU
      2 STU
      9 STU
      s1 s3 XCHG
      256 STU
      s0 s2 XCHG
      <{
        ONE
        ROT
        98 STU
        STREF
      }> PUSHCONT
      <{
        ZERO
        ROT
        98 STU
        SWAP
        CTOS
        STSLICER
      }> PUSHCONT
      IFELSE
      ENDC
      16 PUSHINT
      SENDRAWMSG
      s4 s2 XCHG2
      SUB
      ROT
      SUB
      14 GETGLOB
      s1 s(-1) PUXC
      GEQ
      101 THROWIFNOT
      16 PUSHINT
      RAWRESERVE
      OVER
      2 PLDU
      ISZERO
      NOT
      <{
        NEWC
        3576854235 PUSHINT
        ROT
        s0 s2 XCHG
        32 STU
        64 STU
        ENDC
        ZERO
        SWAP
        16 PUSHINT
        130 PUSHINT
        s2 PUSH
        ISNULL
        NOT
        DUP
        <{
          ONE
        }> PUSHCONT
        <{
          ZERO
        }> PUSHCONT
        IFELSE
        NEWC
        s0 s4 XCHG2
        SWAP
        6 STU
        s0 s6 XCHG2
        STSLICER
        s0 s4 XCHG2
        STGRAMS
        107 STU
        s0 s3 XCHG
        <{
          ROT
          STREF
          SWAP
        }> PUSHCONT
        <{
          DROP
        }> PUSHCONT
        IFELSE
        SWAP
        ENDC
        SWAP
        SENDRAWMSG
      }> PUSHCONT
      <{
        DROP2
      }> PUSHCONT
      IFELSE
    }>c IFJMPREF
    DROP
    s5 POP
    OVER
    3450270305 PUSHINT
    EQUAL
    <{
      NIP
      s3 POP
      s4 POP
      s0 s3 XCHG
      256 LDU
      SWAP
      MYADDR
      17 GETGLOB
      s2 s(-1) PUXC
      s0 s2 XCHG
      ZERO
      ZERO
      s0 s3 XCHG2
      ZERO
      NEWC
      s0 s4 XCHG2
      STSLICER
      x{2_} PUSHSLICE
      STSLICER
      s1 s2 XCHG
      1 STI
      SWAP
      SWAP
      256 STU
      1 STU
      ENDC
      OVER
      NEWC
      2 STU
      s1 s3 XCHG
      STDICT
      s1 s2 XCHG
      STDICT
      1 STU
      ENDC
      ZERO
      SWAP
      HASHCU
      4 PUSHINT
      NEWC
      3 STU
      s1 s2 XCHG
      8 STI
      256 STU
      ENDC
      CTOS
      s0 s5 XCHG2
      SDEQ
      71 THROWIFNOT
      s0 s2 XCHG
      20 PUSHINT
      RAWRESERVE
      SWAP
      LDMSGADDR
      DROP
      NEWC
      s3 PUSH
      SWAP
      256 STU
      ENDC
      3450270305 PUSHINT
      SWAP
      -1 PUSHINT
      8 PUSHPOW2
      ONE
      12 PUSHINT
      NEWC
      4 STU
      2 STU
      9 STU
      s1 s3 XCHG
      256 STU
      s0 s2 XCHG
      <{
        ONE
        ROT
        98 STU
        STREF
      }> PUSHCONT
      <{
        ZERO
        ROT
        98 STU
        SWAP
        CTOS
        STSLICER
      }> PUSHCONT
      IFELSE
      ENDC
      16 PUSHINT
      SENDRAWMSG
      DUP
      2 PLDU
      ISZERO
      NOT
      <{
        NEWC
        2264036229 PUSHINT
        ROT
        s0 s3 XCHG
        s0 s2 XCHG
        32 STU
        64 STU
        ROT
        SWAP
        256 STU
        ENDC
        ZERO
        SWAP
        16 PUSHINT
        130 PUSHINT
        s2 PUSH
        ISNULL
        NOT
        DUP
        <{
          ONE
        }> PUSHCONT
        <{
          ZERO
        }> PUSHCONT
        IFELSE
        NEWC
        s0 s4 XCHG2
        SWAP
        6 STU
        s0 s6 XCHG2
        STSLICER
        s0 s4 XCHG2
        STGRAMS
        107 STU
        s0 s3 XCHG
        <{
          ROT
          STREF
          SWAP
        }> PUSHCONT
        <{
          DROP
        }> PUSHCONT
        IFELSE
        SWAP
        ENDC
        SWAP
        SENDRAWMSG
      }> PUSHCONT
      <{
        3 BLKDROP
      }> PUSHCONT
      IFELSE
      NEWC
      15 GETGLOB
      STSLICER
      16 GETGLOB
      STGRAMS
      ENDC
      6 GETGLOB
      5 GETGLOB
      NEWC
      11 GETGLOB
      SWAP
      48 STU
      7 GETGLOB
      SWAP
      48 STU
      8 GETGLOB
      SWAP
      48 STU
      9 GETGLOB
      SWAP
      48 STU
      10 GETGLOB
      SWAP
      48 STU
      4 STU
      STDICT
      ENDC
      17 GETGLOB
      NEWC
      4 GETGLOB
      SWAP
      256 STU
      STREF
      ENDC
      NEWC
      13 GETGLOB
      STGRAMS
      12 GETGLOB
      STGRAMS
      14 GETGLOB
      STGRAMS
      ENDC
      NEWC
      1 GETGLOB
      STSLICER
      2 GETGLOB
      STSLICER
      3 GETGLOB
      STSLICER
      STREF
      STREF
      STREF
      STREF
      ENDC
      POPROOT
    }> PUSHCONT
    IFJMP
    OVER
    1959622322 PUSHINT
    EQUAL
    <{
      DROP2
      s2 POP
      s3 POP
      3 GETGLOB
      s1 s3 XCHG
      SDEQ
      72 THROWIFNOT
      SWAP
      256 LDU
      SWAP
      NIP
      4 SETGLOB
      20 PUSHINT
      RAWRESERVE
      NEWC
      15 GETGLOB
      STSLICER
      16 GETGLOB
      STGRAMS
      ENDC
      6 GETGLOB
      5 GETGLOB
      NEWC
      11 GETGLOB
      SWAP
      48 STU
      7 GETGLOB
      SWAP
      48 STU
      8 GETGLOB
      SWAP
      48 STU
      9 GETGLOB
      SWAP
      48 STU
      10 GETGLOB
      SWAP
      48 STU
      4 STU
      STDICT
      ENDC
      17 GETGLOB
      NEWC
      4 GETGLOB
      SWAP
      256 STU
      STREF
      ENDC
      NEWC
      13 GETGLOB
      STGRAMS
      12 GETGLOB
      STGRAMS
      14 GETGLOB
      STGRAMS
      ENDC
      NEWC
      1 GETGLOB
      STSLICER
      2 GETGLOB
      STSLICER
      3 GETGLOB
      STSLICER
      STREF
      STREF
      STREF
      STREF
      ENDC
      POPROOT
    }>c IFJMPREF
    s5 POP
    DUP
    529922156 PUSHINT
    EQUAL
    <{
      s4 s5 XCHG
      5 BLKDROP
      14 GETGLOB
      0 NEQINT
      200 THROWIFNOT
      NEWC
      4082677456 PUSHINT
      ROT
      s0 s2 XCHG
      32 STU
      64 STU
      14 GETGLOB
      STGRAMS
      ENDC
      1 GETGLOB
      14 GETGLOB
      ROT
      16 PUSHINT
      TWO
      s2 PUSH
      ISNULL
      NOT
      DUP
      <{
        ONE
      }> PUSHCONT
      <{
        ZERO
      }> PUSHCONT
      IFELSE
      NEWC
      s0 s4 XCHG2
      SWAP
      6 STU
      s0 s6 XCHG2
      STSLICER
      s0 s4 XCHG2
      STGRAMS
      107 STU
      s0 s3 XCHG
      <{
        ROT
        STREF
        SWAP
      }> PUSHCONT
      <{
        DROP
      }> PUSHCONT
      IFELSE
      SWAP
      ENDC
      SWAP
      SENDRAWMSG
      ZERO
      14 SETGLOB
      NEWC
      15 GETGLOB
      STSLICER
      16 GETGLOB
      STGRAMS
      ENDC
      6 GETGLOB
      5 GETGLOB
      NEWC
      11 GETGLOB
      SWAP
      48 STU
      7 GETGLOB
      SWAP
      48 STU
      8 GETGLOB
      SWAP
      48 STU
      9 GETGLOB
      SWAP
      48 STU
      10 GETGLOB
      SWAP
      48 STU
      4 STU
      STDICT
      ENDC
      17 GETGLOB
      NEWC
      4 GETGLOB
      SWAP
      256 STU
      STREF
      ENDC
      NEWC
      13 GETGLOB
      STGRAMS
      12 GETGLOB
      STGRAMS
      14 GETGLOB
      STGRAMS
      ENDC
      NEWC
      1 GETGLOB
      STSLICER
      2 GETGLOB
      STSLICER
      3 GETGLOB
      STSLICER
      STREF
      STREF
      STREF
      STREF
      ENDC
      POPROOT
    }>c IFJMPREF
    DUP
    103913911 PUSHINT
    EQUAL
    <{
      DROP
      2 2 BLKDROP2
      s2 POP
      1 GETGLOB
      SDEQ
      70 THROWIFNOT
      LDGRAMS
      DROP
      12 SETGLOB
      NEWC
      15 GETGLOB
      STSLICER
      16 GETGLOB
      STGRAMS
      ENDC
      6 GETGLOB
      5 GETGLOB
      NEWC
      11 GETGLOB
      SWAP
      48 STU
      7 GETGLOB
      SWAP
      48 STU
      8 GETGLOB
      SWAP
      48 STU
      9 GETGLOB
      SWAP
      48 STU
      10 GETGLOB
      SWAP
      48 STU
      4 STU
      STDICT
      ENDC
      17 GETGLOB
      NEWC
      4 GETGLOB
      SWAP
      256 STU
      STREF
      ENDC
      NEWC
      13 GETGLOB
      STGRAMS
      12 GETGLOB
      STGRAMS
      14 GETGLOB
      STGRAMS
      ENDC
      NEWC
      1 GETGLOB
      STSLICER
      2 GETGLOB
      STSLICER
      3 GETGLOB
      STSLICER
      STREF
      STREF
      STREF
      STREF
      ENDC
      POPROOT
    }> PUSHCONT
    IFJMP
    DUP
    892421724 PUSHINT
    EQUAL
    <{
      DROP
      2 2 BLKDROP2
      s2 POP
      1 GETGLOB
      SDEQ
      70 THROWIFNOT
      LDGRAMS
      DROP
      13 SETGLOB
      NEWC
      15 GETGLOB
      STSLICER
      16 GETGLOB
      STGRAMS
      ENDC
      6 GETGLOB
      5 GETGLOB
      NEWC
      11 GETGLOB
      SWAP
      48 STU
      7 GETGLOB
      SWAP
      48 STU
      8 GETGLOB
      SWAP
      48 STU
      9 GETGLOB
      SWAP
      48 STU
      10 GETGLOB
      SWAP
      48 STU
      4 STU
      STDICT
      ENDC
      17 GETGLOB
      NEWC
      4 GETGLOB
      SWAP
      256 STU
      STREF
      ENDC
      NEWC
      13 GETGLOB
      STGRAMS
      12 GETGLOB
      STGRAMS
      14 GETGLOB
      STGRAMS
      ENDC
      NEWC
      1 GETGLOB
      STSLICER
      2 GETGLOB
      STSLICER
      3 GETGLOB
      STSLICER
      STREF
      STREF
      STREF
      STREF
      ENDC
      POPROOT
    }>c IFJMPREF
    DUP
    1558997984 PUSHINT
    EQUAL
    <{
      DROP
      2 2 BLKDROP2
      s2 POP
      1 GETGLOB
      SDEQ
      70 THROWIFNOT
      LDMSGADDR
      DROP
      3 SETGLOB
      NEWC
      15 GETGLOB
      STSLICER
      16 GETGLOB
      STGRAMS
      ENDC
      6 GETGLOB
      5 GETGLOB
      NEWC
      11 GETGLOB
      SWAP
      48 STU
      7 GETGLOB
      SWAP
      48 STU
      8 GETGLOB
      SWAP
      48 STU
      9 GETGLOB
      SWAP
      48 STU
      10 GETGLOB
      SWAP
      48 STU
      4 STU
      STDICT
      ENDC
      17 GETGLOB
      NEWC
      4 GETGLOB
      SWAP
      256 STU
      STREF
      ENDC
      NEWC
      13 GETGLOB
      STGRAMS
      12 GETGLOB
      STGRAMS
      14 GETGLOB
      STGRAMS
      ENDC
      NEWC
      1 GETGLOB
      STSLICER
      2 GETGLOB
      STSLICER
      3 GETGLOB
      STSLICER
      STREF
      STREF
      STREF
      STREF
      ENDC
      POPROOT
    }> PUSHCONT
    IFJMP
    DUP
    1477999036 PUSHINT
    EQUAL
    <{
      DROP
      2 2 BLKDROP2
      s2 POP
      1 GETGLOB
      SDEQ
      70 THROWIFNOT
      LDMSGADDR
      DROP
      2 SETGLOB
      2 GETGLOB
      x{2_} PUSHSLICE
      SDEQ
      80 THROWIF
      NEWC
      15 GETGLOB
      STSLICER
      16 GETGLOB
      STGRAMS
      ENDC
      6 GETGLOB
      5 GETGLOB
      NEWC
      11 GETGLOB
      SWAP
      48 STU
      7 GETGLOB
      SWAP
      48 STU
      8 GETGLOB
      SWAP
      48 STU
      9 GETGLOB
      SWAP
      48 STU
      10 GETGLOB
      SWAP
      48 STU
      4 STU
      STDICT
      ENDC
      17 GETGLOB
      NEWC
      4 GETGLOB
      SWAP
      256 STU
      STREF
      ENDC
      NEWC
      13 GETGLOB
      STGRAMS
      12 GETGLOB
      STGRAMS
      14 GETGLOB
      STGRAMS
      ENDC
      NEWC
      1 GETGLOB
      STSLICER
      2 GETGLOB
      STSLICER
      3 GETGLOB
      STSLICER
      STREF
      STREF
      STREF
      STREF
      ENDC
      POPROOT
    }>c IFJMPREF
    DUP
    1611221531 PUSHINT
    EQUAL
    <{
      s2 s5 XCHG
      5 BLKDROP
      1 GETGLOB
      SDEQ
      70 THROWIFNOT
      x{2_} PUSHSLICE
      2 SETGLOB
      NEWC
      15 GETGLOB
      STSLICER
      16 GETGLOB
      STGRAMS
      ENDC
      6 GETGLOB
      5 GETGLOB
      NEWC
      11 GETGLOB
      SWAP
      48 STU
      7 GETGLOB
      SWAP
      48 STU
      8 GETGLOB
      SWAP
      48 STU
      9 GETGLOB
      SWAP
      48 STU
      10 GETGLOB
      SWAP
      48 STU
      4 STU
      STDICT
      ENDC
      17 GETGLOB
      NEWC
      4 GETGLOB
      SWAP
      256 STU
      STREF
      ENDC
      NEWC
      13 GETGLOB
      STGRAMS
      12 GETGLOB
      STGRAMS
      14 GETGLOB
      STGRAMS
      ENDC
      NEWC
      1 GETGLOB
      STSLICER
      2 GETGLOB
      STSLICER
      3 GETGLOB
      STSLICER
      STREF
      STREF
      STREF
      STREF
      ENDC
      POPROOT
    }> PUSHCONT
    IFJMP
    DUP
    1783610932 PUSHINT
    EQUAL
    <{
      s2 s5 XCHG
      5 BLKDROP
      2 GETGLOB
      SDEQ
      73 THROWIFNOT
      2 GETGLOB
      1 SETGLOB
      x{2_} PUSHSLICE
      2 SETGLOB
      NEWC
      15 GETGLOB
      STSLICER
      16 GETGLOB
      STGRAMS
      ENDC
      6 GETGLOB
      5 GETGLOB
      NEWC
      11 GETGLOB
      SWAP
      48 STU
      7 GETGLOB
      SWAP
      48 STU
      8 GETGLOB
      SWAP
      48 STU
      9 GETGLOB
      SWAP
      48 STU
      10 GETGLOB
      SWAP
      48 STU
      4 STU
      STDICT
      ENDC
      17 GETGLOB
      NEWC
      4 GETGLOB
      SWAP
      256 STU
      STREF
      ENDC
      NEWC
      13 GETGLOB
      STGRAMS
      12 GETGLOB
      STGRAMS
      14 GETGLOB
      STGRAMS
      ENDC
      NEWC
      1 GETGLOB
      STSLICER
      2 GETGLOB
      STSLICER
      3 GETGLOB
      STSLICER
      STREF
      STREF
      STREF
      STREF
      ENDC
      POPROOT
    }>c IFJMPREF
    DUP
    1852917241 PUSHINT
    EQUAL
    <{
      DROP
      2 2 BLKDROP2
      s2 POP
      1 GETGLOB
      SDEQ
      70 THROWIFNOT
      LDREF
      SWAP
      SETCODE
      LDREF
      DROP
      POPROOT
    }>c IFJMPREF
    DUP
    2128979664 PUSHINT
    EQUAL
    <{
      DROP
      2 2 BLKDROP2
      s2 POP
      1 GETGLOB
      SDEQ
      70 THROWIFNOT
      LDREF
      DROP
      17 SETGLOB
      NULL
      6 SETGLOB
      2128979664 PUSHINT
      NEWC
      ENDC
      ZERO
      8 PUSHPOW2
      ONE
      12 PUSHINT
      NEWC
      4 STU
      2 STU
      9 STU
      s1 s3 XCHG
      256 STU
      s0 s2 XCHG
      <{
        ONE
        ROT
        98 STU
        STREF
      }> PUSHCONT
      <{
        ZERO
        ROT
        98 STU
        SWAP
        CTOS
        STSLICER
      }> PUSHCONT
      IFELSE
      ENDC
      16 PUSHINT
      SENDRAWMSG
      NEWC
      15 GETGLOB
      STSLICER
      16 GETGLOB
      STGRAMS
      ENDC
      6 GETGLOB
      5 GETGLOB
      NEWC
      11 GETGLOB
      SWAP
      48 STU
      7 GETGLOB
      SWAP
      48 STU
      8 GETGLOB
      SWAP
      48 STU
      9 GETGLOB
      SWAP
      48 STU
      10 GETGLOB
      SWAP
      48 STU
      4 STU
      STDICT
      ENDC
      17 GETGLOB
      NEWC
      4 GETGLOB
      SWAP
      256 STU
      STREF
      ENDC
      NEWC
      13 GETGLOB
      STGRAMS
      12 GETGLOB
      STGRAMS
      14 GETGLOB
      STGRAMS
      ENDC
      NEWC
      1 GETGLOB
      STSLICER
      2 GETGLOB
      STSLICER
      3 GETGLOB
      STSLICER
      STREF
      STREF
      STREF
      STREF
      ENDC
      POPROOT
    }> PUSHCONT
    IFJMP
    DUP
    3916583095 PUSHINT
    EQUAL
    <{
      DROP
      2 2 BLKDROP2
      s2 POP
      1 GETGLOB
      SDEQ
      70 THROWIFNOT
      48 LDU
      SWAP
      NIP
      11 SETGLOB
      NEWC
      15 GETGLOB
      STSLICER
      16 GETGLOB
      STGRAMS
      ENDC
      6 GETGLOB
      5 GETGLOB
      NEWC
      11 GETGLOB
      SWAP
      48 STU
      7 GETGLOB
      SWAP
      48 STU
      8 GETGLOB
      SWAP
      48 STU
      9 GETGLOB
      SWAP
      48 STU
      10 GETGLOB
      SWAP
      48 STU
      4 STU
      STDICT
      ENDC
      17 GETGLOB
      NEWC
      4 GETGLOB
      SWAP
      256 STU
      STREF
      ENDC
      NEWC
      13 GETGLOB
      STGRAMS
      12 GETGLOB
      STGRAMS
      14 GETGLOB
      STGRAMS
      ENDC
      NEWC
      1 GETGLOB
      STSLICER
      2 GETGLOB
      STSLICER
      3 GETGLOB
      STSLICER
      STREF
      STREF
      STREF
      STREF
      ENDC
      POPROOT
    }>c IFJMPREF
    DUP
    3576854235 PUSHINT
    EQUAL
    <{
      6 BLKDROP
    }> PUSHCONT
    IFJMP
    DUP
    1960538045 PUSHINT
    EQUAL
    <{
      DROP
      2 2 BLKDROP2
      s2 POP
      1 GETGLOB
      SDEQ
      70 THROWIFNOT
      LDMSGADDR
      DROP
      15 SETGLOB
      NEWC
      15 GETGLOB
      STSLICER
      16 GETGLOB
      STGRAMS
      ENDC
      6 GETGLOB
      5 GETGLOB
      NEWC
      11 GETGLOB
      SWAP
      48 STU
      7 GETGLOB
      SWAP
      48 STU
      8 GETGLOB
      SWAP
      48 STU
      9 GETGLOB
      SWAP
      48 STU
      10 GETGLOB
      SWAP
      48 STU
      4 STU
      STDICT
      ENDC
      17 GETGLOB
      NEWC
      4 GETGLOB
      SWAP
      256 STU
      STREF
      ENDC
      NEWC
      13 GETGLOB
      STGRAMS
      12 GETGLOB
      STGRAMS
      14 GETGLOB
      STGRAMS
      ENDC
      NEWC
      1 GETGLOB
      STSLICER
      2 GETGLOB
      STSLICER
      3 GETGLOB
      STSLICER
      STREF
      STREF
      STREF
      STREF
      ENDC
      POPROOT
    }>c IFJMPREF
    DUP
    345799315 PUSHINT
    EQUAL
    <{
      DROP
      1 GETGLOB
      s1 s2 XCHG
      SDEQ
      70 THROWIFNOT
      LDGRAMS
      DROP
      s3 s3 XCHG2
      SUB
      s2 s(-1) PUXC
      LEQ
      207 THROWIFNOT
      NEWC
      1490661915 PUSHINT
      ROT
      s0 s2 XCHG
      32 STU
      64 STU
      OVER
      STGRAMS
      ENDC
      15 GETGLOB
      s2 s1 PUXC
      24 PUSHINT
      16 PUSHINT
      s2 PUSH
      ISNULL
      NOT
      DUP
      <{
        ONE
      }> PUSHCONT
      <{
        ZERO
      }> PUSHCONT
      IFELSE
      NEWC
      s0 s4 XCHG2
      SWAP
      6 STU
      s0 s6 XCHG2
      STSLICER
      s0 s4 XCHG2
      STGRAMS
      107 STU
      s0 s3 XCHG
      <{
        ROT
        STREF
        SWAP
      }> PUSHCONT
      <{
        DROP
      }> PUSHCONT
      IFELSE
      SWAP
      ENDC
      SWAP
      SENDRAWMSG
      16 GETGLOB
      SWAP
      ADD
      16 SETGLOB
      NEWC
      15 GETGLOB
      STSLICER
      16 GETGLOB
      STGRAMS
      ENDC
      6 GETGLOB
      5 GETGLOB
      NEWC
      11 GETGLOB
      SWAP
      48 STU
      7 GETGLOB
      SWAP
      48 STU
      8 GETGLOB
      SWAP
      48 STU
      9 GETGLOB
      SWAP
      48 STU
      10 GETGLOB
      SWAP
      48 STU
      4 STU
      STDICT
      ENDC
      17 GETGLOB
      NEWC
      4 GETGLOB
      SWAP
      256 STU
      STREF
      ENDC
      NEWC
      13 GETGLOB
      STGRAMS
      12 GETGLOB
      STGRAMS
      14 GETGLOB
      STGRAMS
      ENDC
      NEWC
      1 GETGLOB
      STSLICER
      2 GETGLOB
      STSLICER
      3 GETGLOB
      STSLICER
      STREF
      STREF
      STREF
      STREF
      ENDC
      POPROOT
    }> PUSHCONT
    IFJMP
    DUP
    1176318008 PUSHINT
    EQUAL
    <{
      DROP
      s2 POP
      s3 POP
      1 GETGLOB
      s1 s3 XCHG
      SDEQ
      70 THROWIFNOT
      SWAP
      LDGRAMS
      DROP
      16 GETGLOB
      s1 s(-1) PUXC
      LEQ
      208 THROWIFNOT
      NEWC
      106521215 PUSHINT
      ROT
      s0 s3 XCHG
      s0 s2 XCHG
      32 STU
      64 STU
      OVER
      STGRAMS
      ENDC
      15 GETGLOB
      ZERO
      ROT
      24 PUSHINT
      80 PUSHINT
      s2 PUSH
      ISNULL
      NOT
      DUP
      <{
        ONE
      }> PUSHCONT
      <{
        ZERO
      }> PUSHCONT
      IFELSE
      NEWC
      s0 s4 XCHG2
      SWAP
      6 STU
      s0 s6 XCHG2
      STSLICER
      s0 s4 XCHG2
      STGRAMS
      107 STU
      s0 s3 XCHG
      <{
        ROT
        STREF
        SWAP
      }> PUSHCONT
      <{
        DROP
      }> PUSHCONT
      IFELSE
      SWAP
      ENDC
      SWAP
      SENDRAWMSG
      16 GETGLOB
      SWAP
      SUB
      16 SETGLOB
      NEWC
      15 GETGLOB
      STSLICER
      16 GETGLOB
      STGRAMS
      ENDC
      6 GETGLOB
      5 GETGLOB
      NEWC
      11 GETGLOB
      SWAP
      48 STU
      7 GETGLOB
      SWAP
      48 STU
      8 GETGLOB
      SWAP
      48 STU
      9 GETGLOB
      SWAP
      48 STU
      10 GETGLOB
      SWAP
      48 STU
      4 STU
      STDICT
      ENDC
      17 GETGLOB
      NEWC
      4 GETGLOB
      SWAP
      256 STU
      STREF
      ENDC
      NEWC
      13 GETGLOB
      STGRAMS
      12 GETGLOB
      STGRAMS
      14 GETGLOB
      STGRAMS
      ENDC
      NEWC
      1 GETGLOB
      STSLICER
      2 GETGLOB
      STSLICER
      3 GETGLOB
      STSLICER
      STREF
      STREF
      STREF
      STREF
      ENDC
      POPROOT
    }>c IFJMPREF
    s4 POP
    s0 s3 XCHG
    1185182057 PUSHINT
    EQUAL
    <{
      15 GETGLOB
      SDEQ
      74 THROWIFNOT
      SWAP
      LDDICT
      DROP
      DUP
      ISNULL
      NOT
      <{
        CTOS
        256 LDU
        SWAP
        SWAP
        LDMSGADDR
        LDGRAMS
        LDREF
        LDDICT
        DROP
        14 GETGLOB
        s1 s7 XCHG
        SUB
        s0 s5 XCHG2
        SUB
        OVER
        GEQ
        <{
          NIP
          2 2 BLKDROP2
          NEWC
          ROT
          SWAP
          256 STU
          SWAP
          STGRAMS
          ENDC
          1239573827 PUSHINT
          SWAP
          ZERO
          8 PUSHPOW2
          ONE
          12 PUSHINT
          NEWC
          4 STU
          2 STU
          9 STU
          s1 s3 XCHG
          256 STU
          s0 s2 XCHG
          <{
            ONE
            ROT
            98 STU
            STREF
          }> PUSHCONT
          <{
            ZERO
            ROT
            98 STU
            SWAP
            CTOS
            STSLICER
          }> PUSHCONT
          IFELSE
          ENDC
          16 PUSHINT
          SENDRAWMSG
          COMMIT
          101 THROW
        }> PUSHCONT
        <{
          s0 s3 s4 XCHG3
          DUP
          ISNULL
          <{
            DROP
            16 PUSHINT
            17 PUSHINT
            s2 PUSH
            ISNULL
            NOT
            DUP
            <{
              ONE
            }> PUSHCONT
            <{
              ZERO
            }> PUSHCONT
            IFELSE
            NEWC
            s0 s4 XCHG2
            SWAP
            6 STU
            s0 s6 XCHG2
            STSLICER
            s0 s4 XCHG2
            STGRAMS
            107 STU
            s0 s3 XCHG
            <{
              ROT
              STREF
              SWAP
            }> PUSHCONT
            <{
              DROP
            }> PUSHCONT
            IFELSE
            SWAP
            ENDC
            SWAP
            SENDRAWMSG
          }> PUSHCONT
          <{
            NEWC
            s0 s2 XCHG
            CTOS
            s1 s2 XCHG
            STSLICER
            16 PUSHINT
            17 PUSHINT
            ?fun_ref_7cbbf82cfca64347 INLINECALLDICT
          }> PUSHCONT
          IFELSE
          NEWC
          SWAP
          SWAP
          256 STU
          ENDC
          1185182057 PUSHINT
          SWAP
          ZERO
          8 PUSHPOW2
          ONE
          12 PUSHINT
          NEWC
          4 STU
          2 STU
          9 STU
          s1 s3 XCHG
          256 STU
          s0 s2 XCHG
          <{
            ONE
            ROT
            98 STU
            STREF
          }> PUSHCONT
          <{
            ZERO
            ROT
            98 STU
            SWAP
            CTOS
            STSLICER
          }> PUSHCONT
          IFELSE
          ENDC
          16 PUSHINT
          SENDRAWMSG
        }>c IFREFELSE
      }> PUSHCONT
      <{
        3 BLKDROP
      }> PUSHCONT
      IFELSE
    }>c IFJMPREF
    4 BLKDROP
    16 PUSHPOW2DEC
    THROWANY
  }>
  ?fun_79721 PROC:<{
    PUSHROOT
    CTOS
    LDMSGADDR
    SWAP
    1 SETGLOB
    LDMSGADDR
    SWAP
    2 SETGLOB
    LDMSGADDR
    SWAP
    3 SETGLOB
    LDREF
    SWAP
    CTOS
    LDGRAMS
    SWAP
    13 SETGLOB
    LDGRAMS
    SWAP
    12 SETGLOB
    LDGRAMS
    DROP
    14 SETGLOB
    LDREF
    SWAP
    CTOS
    256 LDU
    SWAP
    4 SETGLOB
    LDREF
    DROP
    17 SETGLOB
    LDREF
    SWAP
    CTOS
    48 LDU
    SWAP
    11 SETGLOB
    48 LDU
    SWAP
    7 SETGLOB
    48 LDU
    SWAP
    8 SETGLOB
    48 LDU
    SWAP
    9 SETGLOB
    48 LDU
    SWAP
    10 SETGLOB
    4 LDU
    SWAP
    5 SETGLOB
    LDDICT
    DROP
    6 SETGLOB
    LDREF
    DROP
    CTOS
    LDMSGADDR
    SWAP
    15 SETGLOB
    LDGRAMS
    DROP
    16 SETGLOB
    15 GETGLOB
    16 GETGLOB
  }>
  ?fun_82469 PROC:<{
    PUSHROOT
    CTOS
    LDMSGADDR
    SWAP
    1 SETGLOB
    LDMSGADDR
    SWAP
    2 SETGLOB
    LDMSGADDR
    SWAP
    3 SETGLOB
    LDREF
    SWAP
    CTOS
    LDGRAMS
    SWAP
    13 SETGLOB
    LDGRAMS
    SWAP
    12 SETGLOB
    LDGRAMS
    DROP
    14 SETGLOB
    LDREF
    SWAP
    CTOS
    256 LDU
    SWAP
    4 SETGLOB
    LDREF
    DROP
    17 SETGLOB
    LDREF
    SWAP
    CTOS
    48 LDU
    SWAP
    11 SETGLOB
    48 LDU
    SWAP
    7 SETGLOB
    48 LDU
    SWAP
    8 SETGLOB
    48 LDU
    SWAP
    9 SETGLOB
    48 LDU
    SWAP
    10 SETGLOB
    4 LDU
    SWAP
    5 SETGLOB
    LDDICT
    DROP
    6 SETGLOB
    LDREF
    DROP
    CTOS
    LDMSGADDR
    SWAP
    15 SETGLOB
    LDGRAMS
    DROP
    16 SETGLOB
    1 GETGLOB
    2 GETGLOB
    3 GETGLOB
    5 GETGLOB
    6 GETGLOB
    7 GETGLOB
    8 GETGLOB
    9 GETGLOB
    11 GETGLOB
    10 GETGLOB
    13 GETGLOB
    12 GETGLOB
    14 GETGLOB
    17 GETGLOB
    4 GETGLOB
    15 GETGLOB
    16 GETGLOB
  }>
  ?fun_91147 PROC:<{
    PUSHROOT
    CTOS
    LDMSGADDR
    SWAP
    1 SETGLOB
    LDMSGADDR
    SWAP
    2 SETGLOB
    LDMSGADDR
    SWAP
    3 SETGLOB
    LDREF
    SWAP
    CTOS
    LDGRAMS
    SWAP
    13 SETGLOB
    LDGRAMS
    SWAP
    12 SETGLOB
    LDGRAMS
    DROP
    14 SETGLOB
    LDREF
    SWAP
    CTOS
    256 LDU
    SWAP
    4 SETGLOB
    LDREF
    DROP
    17 SETGLOB
    LDREF
    SWAP
    CTOS
    48 LDU
    SWAP
    11 SETGLOB
    48 LDU
    SWAP
    7 SETGLOB
    48 LDU
    SWAP
    8 SETGLOB
    48 LDU
    SWAP
    9 SETGLOB
    48 LDU
    SWAP
    10 SETGLOB
    4 LDU
    SWAP
    5 SETGLOB
    LDDICT
    DROP
    6 SETGLOB
    LDREF
    DROP
    CTOS
    LDMSGADDR
    SWAP
    15 SETGLOB
    LDGRAMS
    DROP
    16 SETGLOB
    ZERO
    DUP
    6 GETGLOB
    8 PUSHPOW2
    DICTUMIN
    NULLSWAPIFNOT2
    <{
    }> PUSHCONT
    <{
      SWAP
      48 LDU
      SWAP
      NIP
      s0 s3 PUSH2
      GREATER
      <{
        2 2 BLKDROP2
        s1 s0 XCPU
      }> PUSHCONT
      <{
        DROP
      }> PUSHCONT
      IFELSE
      6 GETGLOB
      8 PUSHPOW2
      DICTUGETNEXT
      NULLSWAPIFNOT2
    }> PUSHCONT
    WHILE
    s2 s3 XCHG
    3 BLKDROP
    7 GETGLOB
    8 GETGLOB
    9 GETGLOB
    10 GETGLOB
    11 GETGLOB
    5 GETGLOB
    4 GETGLOB
  }>
  ?fun_104069 PROC:<{
    PUSHROOT
    CTOS
    LDMSGADDR
    SWAP
    1 SETGLOB
    LDMSGADDR
    SWAP
    2 SETGLOB
    LDMSGADDR
    SWAP
    3 SETGLOB
    LDREF
    SWAP
    CTOS
    LDGRAMS
    SWAP
    13 SETGLOB
    LDGRAMS
    SWAP
    12 SETGLOB
    LDGRAMS
    DROP
    14 SETGLOB
    LDREF
    SWAP
    CTOS
    256 LDU
    SWAP
    4 SETGLOB
    LDREF
    DROP
    17 SETGLOB
    LDREF
    SWAP
    CTOS
    48 LDU
    SWAP
    11 SETGLOB
    48 LDU
    SWAP
    7 SETGLOB
    48 LDU
    SWAP
    8 SETGLOB
    48 LDU
    SWAP
    9 SETGLOB
    48 LDU
    SWAP
    10 SETGLOB
    4 LDU
    SWAP
    5 SETGLOB
    LDDICT
    DROP
    6 SETGLOB
    LDREF
    DROP
    CTOS
    LDMSGADDR
    SWAP
    15 SETGLOB
    LDGRAMS
    DROP
    16 SETGLOB
    MYADDR
    17 GETGLOB
    s1 s2 XCHG
    s0 s2 XCHG
    ZERO
    ZERO
    s0 s3 XCHG2
    ZERO
    NEWC
    s0 s4 XCHG2
    STSLICER
    x{2_} PUSHSLICE
    STSLICER
    s1 s2 XCHG
    1 STI
    SWAP
    SWAP
    256 STU
    1 STU
    ENDC
    OVER
    NEWC
    2 STU
    s1 s3 XCHG
    STDICT
    s1 s2 XCHG
    STDICT
    1 STU
    ENDC
    DUP
    ZERO
    SWAP
    HASHCU
    4 PUSHINT
    NEWC
    3 STU
    s1 s2 XCHG
    8 STI
    256 STU
    ENDC
    CTOS
    SWAP
  }>
  ?fun_114059 PROC:<{
    GASCONSUMED
    SWAP
    13 PUSHPOW2
    CDATASIZE
    DROP
    GASCONSUMED
    s0 s3 XCHG2
    SUB
    SWAP
    105 ADDCONST
    51611 PUSHINT
    s0 s3 XCHG2
    ADD
    s1 s2 XCHG
    31536000 PUSHINT
    ZERO
    GETSTOR.com/cli/cli/v2/api"
	"github.com/cli/cli/v2/internal/ghrepo"
	"github.com/cli/cli/v2/pkg/iostreams"
)

// Cap the number of git remotes to look up, since the user might have an
// unusually large number of git remotes.
const defaultRemotesForLookup = 5

func ResolveRemotesToRepos(remotes Remotes, client *api.Client, base string) (*ResolvedRemotes, error) {
	sort.Stable(remotes)

	result := &ResolvedRemotes{
		remotes:   remotes,
		apiClient: client,
	}

	var baseOverride ghrepo.Interface
	if base != "" {
		var err error
		baseOverride, err = ghrepo.FromFullName(base)
		if err != nil {
			return result, err
		}
		result.baseOverride = baseOverride
	}

	return result, nil
}

func resolveNetwork(result *ResolvedRemotes, remotesForLookup int) error {
	var repos []ghrepo.Interface
	for _, r := range result.remotes {
		repos = append(repos, r)
		if len(repos) == remotesForLookup {
			break
		}
	}

	networkResult, err := api.RepoNetwork(result.apiClient, repos)
	result.network = &networkResult
	return err
}

type ResolvedRemotes struct {
	baseOverride ghrepo.Interface
	remotes      Remotes
	network      *api.RepoNetworkResult
	apiClient    *api.Client
}

func (r *ResolvedRemotes) BaseRepo(io *iostreams.IOStreams) (ghrepo.Interface, error) {
	if r.baseOverride != nil {
		return r.baseOverride, nil
	}

	if len(r.remotes) == 0 {
		return nil, errors.New("no git remotes")
	}

	// if any of the remotes already has a resolution, respect that
	for _, r := range r.remotes {
		if r.Resolved == "base" {
			return r, nil
		} else if r.Resolved != "" {
			repo, err := ghrepo.FromFullName(r.Resolved)
			if err != nil {
				return nil, err
			}
			return ghrepo.NewWithHost(repo.RepoOwner(), repo.RepoName(), r.RepoHost()), nil
		}
	}

	if !io.CanPrompt() {
		// we cannot prompt, so just resort to the 1st remote
		return r.remotes[0], nil
	}

	repos, err := r.NetworkRepos(defaultRemotesForLookup)
	if err != nil {
		return nil, err
	}

	if len(repos) == 0 {
		return r.remotes[0], nil
	} else if len(repos) == 1 {
		return repos[0], nil
	}

	cs := io.ColorScheme()

	fmt.Fprintf(io.ErrOut,
		"%s No default remote repository has been set. To learn more about the default repository, run: gh repo set-default --help\n",
		cs.FailureIcon())

	fmt.Fprintln(io.Out)

	return nil, errors.New(
		"please run `gh repo set-default` to select a default remote repository.")
}

func (r *ResolvedRemotes) HeadRepos() ([]*api.Repository, error) {
	if r.network == nil {
		err := resolveNetwork(r, defaultRemotesForLookup)
		if err != nil {
			return nil, err
		}
	}

	var results []*api.Repository
	var ids []string // Check if repo duplicates
	for _, repo := range r.network.Repositories {
		if repo != nil && repo.ViewerCanPush() && !slices.Contains(ids, repo.ID) {
			results = append(results, repo)
			ids = append(ids, repo.ID)
		}
	}
	return results, nil
}

// NetworkRepos fetches info about remotes for the network of repos.
// Pass a value of 0 to fetch info on all remotes.
func (r *ResolvedRemotes) NetworkRepos(remotesForLookup int) ([]*api.Repository, error) {
	if r.network == nil {
		err := resolveNetwork(r, remotesForLookup)
		if err != nil {
			return nil, err
		}
	}

	var repos []*api.Repository
	repoMap := map[string]bool{}

	add := func(r *api.Repository) {
		fn := ghrepo.FullName(r)
		if _, ok := repoMap[fn]; !ok {
			repoMap[fn] = true
			repos = append(repos, r)
		}
	}

	for _, repo := range r.network.Repositories {
		if repo == nil {
			continue
		}
		if repo.Parent != nil {
			add(repo.Parent)
		}
		add(repo)
	}

	return repos, nil
}

// RemoteForRepo finds the git remote that points to a repository
func (r *ResolvedRemotes) RemoteForRepo(repo ghrepo.Interface) (*Remote, error) {
	for _, remote := range r.remotes {
		if ghrepo.IsSame(remote, repo) {
			return remote, nil
		}
	}
	return nil, errors.New("not found")
}
