// Copyright (c) Rspscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

import React from "react"
import RspscaleIcon from "src/assets/icons/rspscale-icon.svg?react"

/**
 * DisconnectedView is rendered after node logout.
 */
export default function DisconnectedView() {
  return (
    <>
      <RspscaleIcon className="mx-auto" />
      <p className="mt-12 text-center text-text-muted">
        You logged out of this device. To reconnect it you will have to
        re-authenticate the device from either the Rspscale app or the
        Rspscale command line interface.
      </p>
    </>
  )
}
