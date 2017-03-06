using Go = import "/go.capnp";
@0xb545ff514b234a75;
$Go.package("rpc");
$Go.import("github.com/griff/thonix/rpc");

enum Status {
  booting @0;
  bootFailed @1;
  installing @2;
  running @3;
}

interface Server {
  # Interface of the public API for ThoNix. While ThoNix daemon is running this
  # interface is exported as a socket at "/var/run/thonix.sock".

  const socketPath :Text = "/var/run/thonix.sock";

  state @0 () -> (state: Status);
}

interface ServerAdmin extends (Server) {
  # Interface of the administrator API for ThoNix. While ThoNix daemon is
  # running this interface is exported as a socket accessible by root at 
  # "/var/run/thonix-admin.sock".

  const socketPath :Text = "/var/run/thonix-admin.sock";

  setState @0 (newState: Status) -> ();
}
