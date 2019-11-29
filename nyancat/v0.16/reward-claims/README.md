# Nyancat Testnet Rewards

## point-to-iris conversion rate

1 point = 70 IRIS

## Credit Confirmation

We just assume that you have done the [HTLC Task](../README.md#htlc-tasks), so the following sheet has default added 30 points to all of the participants. **But the [HTLC Task](../README.md#htlc-tasks) is optional for claiming the rewards, you can decide not to do this task, then we can send you the OTHER TASK REWARDS(except HTLC) directly.**

Nyancat Testnet Points Summary: <https://docs.google.com/spreadsheets/d/1a1xvoytSt9zagGecReS9goa7obU302Q0xs97AwCGdCE/edit?usp=sharing>

*If you have any questions, please contact our staff*

## Submit Signature

You will need to provide an IRISHub Mainnet address to receive Nyancat Testnet rewards.

### Signing

Please sign your Mainnet address using the `Keybase PGP` you submitted on the [Nyancat Testnet Registration Form](http://nyancat-irisnet.mikecrm.com/SnqhRqw).

Choose one of the following two signature methods:

- Sign with the [Online Tool](https://keybase.io/sign)

- Or sign with the [Keybase Command Line Tool](https://keybase.io/docs/command_line)

    ```bash
    keybase pgp sign -m "Your IRISHub Mainnet Address"
    ```

### Submitting

You need to save the above signed content in `[github-user-name]-[keybase fingerprint].txt`

For example: `irisnetvalidator-6763B2C7947A9363.txt`

Then use the Github account you submitted on the [Nyancat Testnet Registration Form](http://nyancat-irisnet.mikecrm.com/SnqhRqw) to `Pull Request` the signature file to [sig-files](https://github.com/irisnet/testnets/tree/master/nyancat/v0.16/reward-claims/sig-files).

Once the team receives your PR, it is deemed that you have completed the confirmation of the points and rewards of the Nyancat Testnet, and will send you the Mainnet reward soon.
