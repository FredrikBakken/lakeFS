"""
    lakeFS API

    lakeFS HTTP API  # noqa: E501

    The version of the OpenAPI document: 0.1.0
    Contact: services@treeverse.io
    Generated by: https://openapi-generator.tech
"""


import unittest

import lakefs_client
from lakefs_client.api.statistics_api import StatisticsApi  # noqa: E501


class TestStatisticsApi(unittest.TestCase):
    """StatisticsApi unit test stubs"""

    def setUp(self):
        self.api = StatisticsApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_post_stats_events(self):
        """Test case for post_stats_events

        post stats events, this endpoint is meant for internal use only  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
