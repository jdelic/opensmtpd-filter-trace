Golang implementation of a trace filter for OpenSMTPD
=====================================================

This is a simple implementation based off my
`opensmtpd-filters-go <osfgo_>`__ library.

It logs all events received from opensmtpd to ``/tmp/filterlog.txt``

Example usage in smtpd.conf
---------------------------

::

    filter "trace" proc-exec "/usr/lib/x86_64-linux-gnu/opensmtpd/filter-trace"
    listen on "127.0.0.1" port 25 filter trace


.. _osfgo: https://github.com/jdelic/opensmtpd-filters-go
